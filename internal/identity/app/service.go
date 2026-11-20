package app

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/raire/raire/internal/contracts"
	"github.com/raire/raire/internal/identity/domain"
	"github.com/raire/raire/internal/platform/audit"
	"github.com/raire/raire/internal/platform/clock"
	"github.com/raire/raire/internal/platform/db"
	"github.com/raire/raire/internal/platform/id"
	"github.com/raire/raire/internal/platform/idempotency"
	"github.com/raire/raire/internal/platform/metrics"
	"github.com/raire/raire/internal/platform/outbox"
	"github.com/raire/raire/internal/platform/security"
)

var ErrNotFound = errors.New("not found")

type Repo interface {
	CreateAccount(ctx context.Context, q db.Querier, a domain.Account) error
	GetAccountByEmail(ctx context.Context, q db.Querier, email string) (*domain.Account, error)
	GetAccountByID(ctx context.Context, q db.Querier, id string) (*domain.Account, error)
	UpdateAccount(ctx context.Context, q db.Querier, a domain.Account) error
	SetAccountRole(ctx context.Context, q db.Querier, id string, role domain.Role) error

	CreateConsent(ctx context.Context, q db.Querier, c domain.Consent) error
	ListConsents(ctx context.Context, q db.Querier, accountID string) ([]domain.Consent, error)
	RevokeConsent(ctx context.Context, q db.Querier, accountID string, kind domain.ConsentKind, at time.Time) error

	CreateEmailToken(ctx context.Context, q db.Querier, t domain.EmailToken) error
	GetEmailTokenByHash(ctx context.Context, q db.Querier, hash string) (*domain.EmailToken, error)
	MarkEmailTokenUsed(ctx context.Context, q db.Querier, id string, at time.Time) (bool, error)
	InvalidateEmailTokens(ctx context.Context, q db.Querier, accountID, kind string, at time.Time) error

	CreateSession(ctx context.Context, q db.Querier, s domain.Session) error
	GetSessionByTokenHash(ctx context.Context, q db.Querier, hash string) (*domain.Session, error)
	RevokeSessionByHash(ctx context.Context, q db.Querier, hash string, at time.Time) error
	RevokeSessionByID(ctx context.Context, q db.Querier, id string, at time.Time) error
	RevokeOtherSessions(ctx context.Context, q db.Querier, accountID, keepSessionID string, at time.Time) error
	RevokeAllSessions(ctx context.Context, q db.Querier, accountID string, at time.Time) error
	TouchSession(ctx context.Context, q db.Querier, id string, at time.Time) error

	UpsertVerification(ctx context.Context, q db.Querier, accountID, kind, status string, at time.Time) error
	ReplaceRecoveryCodes(ctx context.Context, q db.Querier, accountID string, hashes []string, at time.Time) error
	ConsumeRecoveryCode(ctx context.Context, q db.Querier, accountID, hash string, at time.Time) (bool, error)
}

type Config struct {
	SessionTTL      time.Duration
	ConfirmTokenTTL time.Duration
	ResetTokenTTL   time.Duration
	TermsVersion    string
	PrivacyVersion  string
	Argon2          security.Params
	Cipher          *security.Cipher
}

type Service struct {
	uow     *db.UoW
	repo    Repo
	audit   *audit.Writer
	outbox  *outbox.Writer
	idem    *idempotency.Store
	clock   clock.Clock
	metrics *metrics.Registry
	cfg     Config
	dummy   string
	cipher  *security.Cipher
}

func NewService(uow *db.UoW, repo Repo, auditWriter *audit.Writer, outboxWriter *outbox.Writer, idem *idempotency.Store, c clock.Clock, m *metrics.Registry, cfg Config) *Service {
	if c == nil {
		c = clock.System{}
	}
	if cfg.SessionTTL <= 0 {
		cfg.SessionTTL = 24 * time.Hour
	}
	if cfg.ConfirmTokenTTL <= 0 {
		cfg.ConfirmTokenTTL = 24 * time.Hour
	}
	if cfg.ResetTokenTTL <= 0 {
		cfg.ResetTokenTTL = time.Hour
	}
	if cfg.TermsVersion == "" {
		cfg.TermsVersion = "2026-11-20"
	}
	if cfg.PrivacyVersion == "" {
		cfg.PrivacyVersion = "2026-11-20"
	}
	if cfg.Argon2.Memory == 0 {
		cfg.Argon2 = security.DefaultParams()
	}
	dummy, _ := security.HashPassword("dummy-password-for-timing-1", cfg.Argon2)
	return &Service{uow: uow, repo: repo, audit: auditWriter, outbox: outboxWriter, idem: idem, clock: c, metrics: m, cfg: cfg, dummy: dummy, cipher: cfg.Cipher}
}

type RequestMeta struct {
	CorrelationID  string
	IP             string
	UserAgent      string
	IdempotencyKey string
	RequestHash    string
}

type SignupInput struct {
	Email         string
	Password      string
	DisplayName   string
	AcceptTerms   bool
	AcceptPrivacy bool
}

type SignupResult struct {
	Account      *domain.Account
	ConfirmToken string
	Created      bool
	Replayed     bool
}

func (s *Service) Signup(ctx context.Context, in SignupInput, meta RequestMeta) (SignupResult, error) {
	email := domain.NormalizeEmail(in.Email)
	if err := domain.ValidateEmail(email); err != nil {
		return SignupResult{}, contracts.Validation("Informe um e-mail válido.")
	}
	if err := domain.ValidatePassword(in.Password); err != nil {
		return SignupResult{}, contracts.Validation(err.Error())
	}
	if err := domain.ValidateDisplayName(in.DisplayName); err != nil {
		return SignupResult{}, contracts.Validation(err.Error())
	}
	if !in.AcceptTerms || !in.AcceptPrivacy {
		return SignupResult{}, contracts.Validation("Aceite os termos e a política de privacidade para continuar.")
	}

	hash, err := security.HashPassword(in.Password, s.cfg.Argon2)
	if err != nil {
		return SignupResult{}, err
	}
	now := s.clock.Now()
	result := SignupResult{}
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.idem.Check(ctx, q, meta.IdempotencyKey, "signup", meta.RequestHash, "", now); err != nil {
			if errors.Is(err, idempotency.ErrReplay) {
				result.Replayed = true
				return nil
			}
			if errors.Is(err, idempotency.ErrConflict) {
				return contracts.New(contracts.CodeConflict, "Esta chave de idempotência já foi usada com outro conteúdo.")
			}
			return err
		}
		if existing, err := s.repo.GetAccountByEmail(ctx, q, email); err == nil && existing != nil {
			result.Created = false
			return nil
		} else if err != nil && !errors.Is(err, ErrNotFound) {
			return err
		}
		acct := domain.Account{
			ID:                id.New("acc"),
			EmailNormalized:   email,
			PasswordHash:      hash,
			DisplayName:       strings.TrimSpace(in.DisplayName),
			Role:              domain.RoleClient,
			ContactEmail:      email,
			ContactPreference: domain.ContactPlatform,
			CreatedAt:         now,
			UpdatedAt:         now,
		}
		if err := s.repo.CreateAccount(ctx, q, acct); err != nil {
			return err
		}
		for _, c := range []struct {
			kind    domain.ConsentKind
			version string
		}{
			{domain.ConsentTerms, s.cfg.TermsVersion},
			{domain.ConsentPrivacy, s.cfg.PrivacyVersion},
		} {
			if err := s.repo.CreateConsent(ctx, q, domain.Consent{
				ID: id.New("cns"), AccountID: acct.ID, Kind: c.kind, Version: c.version, GrantedAt: now,
			}); err != nil {
				return err
			}
		}
		if err := s.repo.UpsertVerification(ctx, q, acct.ID, "contact_email", "not_started", now); err != nil {
			return err
		}
		token, err := s.issueToken(ctx, q, acct.ID, "confirm", s.cfg.ConfirmTokenTTL, now)
		if err != nil {
			return err
		}
		result.Account = &acct
		result.ConfirmToken = token
		result.Created = true
		return s.outbox.Publish(ctx, q, "identity.account_created", map[string]any{
			"account_id": acct.ID, "occurred_at": now,
		})
	})
	if err != nil {
		return SignupResult{}, err
	}
	if result.Created {
		s.count("signup_completed")
	} else {
		s.count("signup_duplicate")
	}
	return result, nil
}

func (s *Service) ConfirmEmail(ctx context.Context, token string, meta RequestMeta) error {
	if strings.TrimSpace(token) == "" {
		return contracts.New(contracts.CodeTokenInvalid, "")
	}
	now := s.clock.Now()
	err := s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.idem.Check(ctx, q, meta.IdempotencyKey, "confirm_email", meta.RequestHash, "", now); err != nil {
			if errors.Is(err, idempotency.ErrReplay) {
				return nil
			}
			return contracts.New(contracts.CodeConflict, "Esta chave de idempotência já foi usada com outro conteúdo.")
		}
		t, err := s.repo.GetEmailTokenByHash(ctx, q, security.HashToken(token))
		if err != nil {
			return contracts.New(contracts.CodeTokenInvalid, "")
		}
		if t.Kind != "confirm" || t.UsedAt != nil {
			return contracts.New(contracts.CodeTokenInvalid, "")
		}
		if !now.Before(t.ExpiresAt) {
			return contracts.New(contracts.CodeTokenExpired, "")
		}
		acct, err := s.repo.GetAccountByID(ctx, q, t.AccountID)
		if err != nil {
			return contracts.New(contracts.CodeTokenInvalid, "")
		}
		used, err := s.repo.MarkEmailTokenUsed(ctx, q, t.ID, now)
		if err != nil {
			return err
		}
		if !used {
			return contracts.New(contracts.CodeTokenInvalid, "")
		}
		acct.EmailVerifiedAt = &now
		acct.UpdatedAt = now
		if err := s.repo.UpdateAccount(ctx, q, *acct); err != nil {
			return err
		}
		if err := s.repo.UpsertVerification(ctx, q, acct.ID, "contact_email", "verified", now); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "identity.email_confirmed", map[string]any{
			"account_id": acct.ID, "occurred_at": now,
		})
	})
	if err == nil {
		s.count("email_confirmed")
	}
	return err
}

func (s *Service) ResendConfirmation(ctx context.Context, email string, meta RequestMeta) (string, error) {
	email = domain.NormalizeEmail(email)
	if err := domain.ValidateEmail(email); err != nil {
		return "", nil // generic response
	}
	now := s.clock.Now()
	var token string
	_ = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.idem.Check(ctx, q, meta.IdempotencyKey, "resend_confirmation", meta.RequestHash, "", now); err != nil {
			if errors.Is(err, idempotency.ErrReplay) {
				return nil
			}
			return err
		}
		acct, err := s.repo.GetAccountByEmail(ctx, q, email)
		if err != nil || acct.EmailVerified() || !acct.Active() {
			return nil
		}
		token, err = s.issueToken(ctx, q, acct.ID, "confirm", s.cfg.ConfirmTokenTTL, now)
		return err
	})
	s.count("confirmation_resent")
	return token, nil
}

type LoginInput struct {
	Email         string
	Password      string
	MFACode       string
	ExistingToken string
}

type LoginResult struct {
	Account      *domain.Account
	SessionToken string
	CSRFToken    string
	ExpiresAt    time.Time
}

func (s *Service) Login(ctx context.Context, in LoginInput, meta RequestMeta) (LoginResult, error) {
	email := domain.NormalizeEmail(in.Email)
	acct, err := s.repo.GetAccountByEmail(ctx, s.uow.Read(), email)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return LoginResult{}, err
	}
	if acct == nil {
		_, _ = security.VerifyPassword(s.dummy, in.Password)
		return LoginResult{}, contracts.New(contracts.CodeUnauthenticated, "E-mail ou senha não conferem.")
	}
	ok, err := security.VerifyPassword(acct.PasswordHash, in.Password)
	if err != nil || !ok {
		s.count("login_failed")
		return LoginResult{}, contracts.New(contracts.CodeUnauthenticated, "E-mail ou senha não conferem.")
	}
	if !acct.Active() {
		return LoginResult{}, contracts.New(contracts.CodeUnauthenticated, "E-mail ou senha não conferem.")
	}
	useRecovery := false
	if acct.Role.Privileged() {
		if strings.TrimSpace(in.MFACode) == "" {
			return LoginResult{}, contracts.New(contracts.CodeMFARequired, "")
		}
		secret, err := s.mfaSecret(acct.MFASecret)
		if err != nil {
			return LoginResult{}, contracts.New(contracts.CodeUnauthenticated, "Código de verificação inválido.")
		}
		if !security.ValidateTOTP(secret, in.MFACode, s.clock.Now()) {
			// Validate and consume the recovery code inside the login transaction
			// so two concurrent logins cannot both use the same code.
			useRecovery = true
		}
	}

	now := s.clock.Now()
	sessionToken, err := security.RandomToken(32)
	if err != nil {
		return LoginResult{}, err
	}
	csrfToken, err := security.RandomToken(32)
	if err != nil {
		return LoginResult{}, err
	}
	sess := domain.Session{
		ID:            id.New("ses"),
		AccountID:     acct.ID,
		TokenHash:     security.HashToken(sessionToken),
		CSRFTokenHash: security.HashToken(csrfToken),
		CreatedAt:     now,
		ExpiresAt:     now.Add(s.cfg.SessionTTL),
		IPHash:        security.Pseudonymize(meta.IP),
		UserAgentHash: security.Pseudonymize(meta.UserAgent),
	}
	if acct.Role.Privileged() {
		sess.MFAVerifiedAt = &now
	}
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if useRecovery {
			ok, err := s.repo.ConsumeRecoveryCode(ctx, q, acct.ID, security.RecoveryCodeHash(in.MFACode), now)
			if err != nil {
				return err
			}
			if !ok {
				s.count("login_mfa_failed")
				return contracts.New(contracts.CodeUnauthenticated, "Código de verificação inválido.")
			}
			if err := s.audit.Write(ctx, q, audit.Event{
				ActorID: acct.ID, ActorRole: string(acct.Role), Action: "account.mfa_recovery_used",
				TargetType: "account", TargetID: acct.ID, CorrelationID: meta.CorrelationID, CreatedAt: now,
			}); err != nil {
				return err
			}
		}
		if in.ExistingToken != "" {
			_ = s.repo.RevokeSessionByHash(ctx, q, security.HashToken(in.ExistingToken), now)
		}
		return s.repo.CreateSession(ctx, q, sess)
	})
	if err != nil {
		return LoginResult{}, err
	}
	s.count("login_succeeded")
	return LoginResult{Account: acct, SessionToken: sessionToken, CSRFToken: csrfToken, ExpiresAt: sess.ExpiresAt}, nil
}

type AuthContext struct {
	Account *domain.Account
	Session *domain.Session
}

func (s *Service) Authenticate(ctx context.Context, token string) (*AuthContext, error) {
	if strings.TrimSpace(token) == "" {
		return nil, contracts.New(contracts.CodeUnauthenticated, "")
	}
	sess, err := s.repo.GetSessionByTokenHash(ctx, s.uow.Read(), security.HashToken(token))
	if err != nil || !sess.Valid(s.clock.Now()) {
		return nil, contracts.New(contracts.CodeUnauthenticated, "")
	}
	acct, err := s.repo.GetAccountByID(ctx, s.uow.Read(), sess.AccountID)
	if err != nil || !acct.Active() {
		return nil, contracts.New(contracts.CodeUnauthenticated, "")
	}
	if acct.Privileged() && sess.MFAVerifiedAt == nil {
		return nil, contracts.New(contracts.CodeMFARequired, "")
	}
	if sess.LastSeenAt == nil || s.clock.Now().Sub(*sess.LastSeenAt) > 5*time.Minute {
		_ = s.repo.TouchSession(ctx, s.uow.Read(), sess.ID, s.clock.Now())
	}
	return &AuthContext{Account: acct, Session: sess}, nil
}

func (s *Service) Logout(ctx context.Context, token string) error {
	if token == "" {
		return nil
	}
	now := s.clock.Now()
	return s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		return s.repo.RevokeSessionByHash(ctx, q, security.HashToken(token), now)
	})
}

func (s *Service) ChangePassword(ctx context.Context, acct *domain.Account, currentSessionID, currentPassword, newPassword string, meta RequestMeta) error {
	ok, err := security.VerifyPassword(acct.PasswordHash, currentPassword)
	if err != nil || !ok {
		return contracts.New(contracts.CodeInvalidPassword, "")
	}
	if err := domain.ValidatePassword(newPassword); err != nil {
		return contracts.Validation(err.Error())
	}
	hash, err := security.HashPassword(newPassword, s.cfg.Argon2)
	if err != nil {
		return err
	}
	now := s.clock.Now()
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		acct.PasswordHash = hash
		acct.UpdatedAt = now
		if err := s.repo.UpdateAccount(ctx, q, *acct); err != nil {
			return err
		}
		if err := s.repo.RevokeOtherSessions(ctx, q, acct.ID, currentSessionID, now); err != nil {
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: acct.ID, ActorRole: string(acct.Role), Action: "account.password_changed",
			TargetType: "account", TargetID: acct.ID, CorrelationID: meta.CorrelationID, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "identity.password_changed", map[string]any{
			"account_id": acct.ID, "occurred_at": now,
		})
	})
	if err == nil {
		s.count("password_changed")
	}
	return err
}

func (s *Service) RequestPasswordReset(ctx context.Context, email string, meta RequestMeta) (string, error) {
	email = domain.NormalizeEmail(email)
	if err := domain.ValidateEmail(email); err != nil {
		return "", nil
	}
	now := s.clock.Now()
	var token string
	_ = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.idem.Check(ctx, q, meta.IdempotencyKey, "password_reset_request", meta.RequestHash, "", now); err != nil {
			if errors.Is(err, idempotency.ErrReplay) {
				return nil
			}
			return err
		}
		acct, err := s.repo.GetAccountByEmail(ctx, q, email)
		if err != nil || !acct.Active() {
			return nil
		}
		token, err = s.issueToken(ctx, q, acct.ID, "reset", s.cfg.ResetTokenTTL, now)
		if err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "identity.password_reset_requested", map[string]any{
			"account_id": acct.ID, "occurred_at": now,
		})
	})
	s.count("password_reset_requested")
	return token, nil
}

func (s *Service) ConfirmPasswordReset(ctx context.Context, token, newPassword string, meta RequestMeta) error {
	if strings.TrimSpace(token) == "" {
		return contracts.New(contracts.CodeTokenInvalid, "")
	}
	if err := domain.ValidatePassword(newPassword); err != nil {
		return contracts.Validation(err.Error())
	}
	hash, err := security.HashPassword(newPassword, s.cfg.Argon2)
	if err != nil {
		return err
	}
	now := s.clock.Now()
	return s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.idem.Check(ctx, q, meta.IdempotencyKey, "password_reset_confirm", meta.RequestHash, "", now); err != nil {
			if errors.Is(err, idempotency.ErrReplay) {
				return nil
			}
			return contracts.New(contracts.CodeConflict, "Esta chave de idempotência já foi usada com outro conteúdo.")
		}
		t, err := s.repo.GetEmailTokenByHash(ctx, q, security.HashToken(token))
		if err != nil {
			return contracts.New(contracts.CodeTokenInvalid, "")
		}
		if t.Kind != "reset" || t.UsedAt != nil {
			return contracts.New(contracts.CodeTokenInvalid, "")
		}
		if !now.Before(t.ExpiresAt) {
			return contracts.New(contracts.CodeTokenExpired, "")
		}
		acct, err := s.repo.GetAccountByID(ctx, q, t.AccountID)
		if err != nil {
			return contracts.New(contracts.CodeTokenInvalid, "")
		}
		used, err := s.repo.MarkEmailTokenUsed(ctx, q, t.ID, now)
		if err != nil {
			return err
		}
		if !used {
			return contracts.New(contracts.CodeTokenInvalid, "")
		}
		acct.PasswordHash = hash
		acct.UpdatedAt = now
		if err := s.repo.UpdateAccount(ctx, q, *acct); err != nil {
			return err
		}
		if err := s.repo.RevokeAllSessions(ctx, q, acct.ID, now); err != nil {
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: acct.ID, ActorRole: string(acct.Role), Action: "account.password_reset_completed",
			TargetType: "account", TargetID: acct.ID, CorrelationID: meta.CorrelationID, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "identity.password_reset_completed", map[string]any{
			"account_id": acct.ID, "occurred_at": now,
		})
	})
}

func (s *Service) issueToken(ctx context.Context, q db.Querier, accountID, kind string, ttl time.Duration, now time.Time) (string, error) {
	token, err := security.RandomToken(32)
	if err != nil {
		return "", err
	}
	if err := s.repo.InvalidateEmailTokens(ctx, q, accountID, kind, now); err != nil {
		return "", err
	}
	if err := s.repo.CreateEmailToken(ctx, q, domain.EmailToken{
		ID: id.New("tok"), AccountID: accountID, Kind: kind,
		TokenHash: security.HashToken(token), ExpiresAt: now.Add(ttl), CreatedAt: now,
	}); err != nil {
		return "", err
	}
	return token, nil
}

func (s *Service) Me(ctx context.Context, accountID string) (*domain.Account, error) {
	acct, err := s.repo.GetAccountByID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	return acct, nil
}

type UpdateMeInput struct {
	DisplayName       *string
	ContactPhone      *string
	ContactPreference *string
}

func (s *Service) UpdateMe(ctx context.Context, acct *domain.Account, in UpdateMeInput, meta RequestMeta) (*domain.Account, error) {
	if in.DisplayName != nil {
		if err := domain.ValidateDisplayName(*in.DisplayName); err != nil {
			return nil, contracts.Validation(err.Error())
		}
		acct.DisplayName = strings.TrimSpace(*in.DisplayName)
	}
	if in.ContactPreference != nil {
		pref := domain.ContactPreference(strings.TrimSpace(*in.ContactPreference))
		if !pref.Valid() {
			return nil, contracts.Validation("Escolha uma preferência de contato válida.")
		}
		acct.ContactPreference = pref
	}
	phoneChanged := false
	if in.ContactPhone != nil {
		phone := strings.TrimSpace(*in.ContactPhone)
		if phone == "" {
			if acct.ContactPhone != "" {
				phoneChanged = true
			}
			acct.ContactPhone = ""
			acct.ContactPhoneVerifiedAt = nil
		} else {
			if err := domain.ValidatePhone(phone); err != nil {
				return nil, contracts.Validation("Informe um telefone válido.")
			}
			normalized := domain.NormalizePhone(phone)
			if normalized != acct.ContactPhone {
				phoneChanged = true
				acct.ContactPhone = normalized
				acct.ContactPhoneVerifiedAt = nil
			}
		}
	}
	now := s.clock.Now()
	acct.UpdatedAt = now
	err := s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.UpdateAccount(ctx, q, *acct); err != nil {
			return err
		}
		if phoneChanged {
			status := "pending"
			if acct.ContactPhone == "" {
				status = "not_started"
			}
			if err := s.repo.UpsertVerification(ctx, q, acct.ID, "contact_phone", status, now); err != nil {
				return err
			}
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: acct.ID, ActorRole: string(acct.Role), Action: "account.updated",
			TargetType: "account", TargetID: acct.ID, CorrelationID: meta.CorrelationID,
			Metadata: map[string]any{"phone_changed": phoneChanged}, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "identity.account_updated", map[string]any{
			"account_id": acct.ID, "phone_changed": phoneChanged, "occurred_at": now,
		})
	})
	return acct, err
}

func (s *Service) ListConsents(ctx context.Context, accountID string) ([]domain.Consent, error) {
	return s.repo.ListConsents(ctx, s.uow.Read(), accountID)
}

func (s *Service) RevokeConsent(ctx context.Context, acct *domain.Account, kind string, meta RequestMeta) error {
	k := domain.ConsentKind(kind)
	if !k.Valid() {
		return contracts.Validation("Consentimento inválido.")
	}
	if k == domain.ConsentTerms || k == domain.ConsentPrivacy {
		return contracts.New(contracts.CodeConflict, "Termos e privacidade são obrigatórios enquanto a conta existir.")
	}
	now := s.clock.Now()
	return s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.RevokeConsent(ctx, q, acct.ID, k, now); err != nil {
			if errors.Is(err, ErrNotFound) {
				return contracts.New(contracts.CodeNotFound, "")
			}
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: acct.ID, ActorRole: string(acct.Role), Action: "account.consent_revoked",
			TargetType: "account", TargetID: acct.ID, CorrelationID: meta.CorrelationID,
			Metadata: map[string]any{"kind": string(k)}, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "identity.consent_changed", map[string]any{
			"account_id": acct.ID, "kind": string(k), "revoked": true, "occurred_at": now,
		})
	})
}

func (s *Service) MarkProfessional(ctx context.Context, accountID string) error {
	acct, err := s.repo.GetAccountByID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return err
	}
	if acct.Role == domain.RoleClient {
		return s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
			return s.repo.SetAccountRole(ctx, q, accountID, domain.RoleProfessional)
		})
	}
	return nil
}

func (s *Service) RevokeSessions(ctx context.Context, accountID string) error {
	return s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		return s.RevokeSessionsTx(ctx, q, accountID)
	})
}

// RevokeSessionsTx lets another context revoke sessions inside the same
// transaction as its own mutation (for example, profile suspension).
func (s *Service) RevokeSessionsTx(ctx context.Context, q db.Querier, accountID string) error {
	return s.repo.RevokeAllSessions(ctx, q, accountID, s.clock.Now())
}

type AccountInfo struct {
	ID                 string
	DisplayName        string
	EmailVerified      bool
	ContactPreference  string
	ContactEmailMasked string
	ContactPhoneMasked string
	PhoneVerified      bool
}

func (s *Service) AccountInfo(ctx context.Context, accountID string) (AccountInfo, error) {
	acct, err := s.repo.GetAccountByID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return AccountInfo{}, err
	}
	info := AccountInfo{
		ID:                 acct.ID,
		DisplayName:        acct.DisplayName,
		EmailVerified:      acct.EmailVerified(),
		ContactPreference:  string(acct.EffectiveContact()),
		ContactEmailMasked: domain.MaskEmail(acct.ContactEmail),
		PhoneVerified:      acct.PhoneVerified(),
	}
	if acct.ContactPhone != "" {
		info.ContactPhoneMasked = domain.MaskPhone(acct.ContactPhone)
	}
	return info, nil
}

// ProvisionMFA generates a TOTP secret and one-time recovery codes and stores
// only the secret and hashes. It is used by the administrative bootstrap, not
// exposed as a public endpoint.
func (s *Service) ProvisionMFA(ctx context.Context, accountID string) (string, []string, error) {
	acct, err := s.repo.GetAccountByID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return "", nil, err
	}
	secret, err := security.GenerateTOTPSecret()
	if err != nil {
		return "", nil, err
	}
	codes := make([]string, 0, 8)
	hashes := make([]string, 0, 8)
	for i := 0; i < 8; i++ {
		code, err := security.RecoveryCode()
		if err != nil {
			return "", nil, err
		}
		codes = append(codes, code)
		hashes = append(hashes, security.RecoveryCodeHash(code))
	}
	acct.MFASecret = secret
	acct.MFARecoveryCodesHash = nil
	acct.UpdatedAt = s.clock.Now()
	if s.cipher == nil {
		return "", nil, errors.New("MFA encryption key is not configured")
	}
	encrypted, err := s.cipher.EncryptString(secret)
	if err != nil {
		return "", nil, err
	}
	acct.MFASecret = "enc:v1:" + encrypted
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.UpdateAccount(ctx, q, *acct); err != nil {
			return err
		}
		return s.repo.ReplaceRecoveryCodes(ctx, q, accountID, hashes, s.clock.Now())
	})
	if err != nil {
		return "", nil, err
	}
	return secret, codes, nil
}

func (s *Service) mfaSecret(stored string) (string, error) {
	if strings.HasPrefix(stored, "enc:v1:") {
		if s.cipher == nil {
			return "", errors.New("MFA encryption key is not configured")
		}
		return s.cipher.DecryptString(strings.TrimPrefix(stored, "enc:v1:"))
	}
	return stored, nil
}

func (s *Service) count(name string) {
	if s.metrics != nil {
		s.metrics.Inc(name)
	}
}

// RequestHash is a helper for handlers implementing Idempotency-Key.
func RequestHash(body []byte) string {
	sum := sha256.Sum256(body)
	return hex.EncodeToString(sum[:])
}

func (s *Service) String() string { return fmt.Sprintf("identity.Service(%s)", s.cfg.TermsVersion) }
