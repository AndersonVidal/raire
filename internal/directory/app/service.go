package app

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/raire/raire/internal/contracts"
	"github.com/raire/raire/internal/directory/domain"
	"github.com/raire/raire/internal/platform/audit"
	"github.com/raire/raire/internal/platform/clock"
	"github.com/raire/raire/internal/platform/db"
	"github.com/raire/raire/internal/platform/id"
	"github.com/raire/raire/internal/platform/metrics"
	"github.com/raire/raire/internal/platform/outbox"
	"github.com/raire/raire/internal/platform/security"
)

var ErrNotFound = errors.New("not found")
var ErrConflict = errors.New("profile version conflict")
var ErrProfileExists = errors.New("account already has a profile")

type AccountInfo struct {
	ID                 string
	DisplayName        string
	EmailVerified      bool
	ContactPreference  string
	ContactEmailMasked string
	ContactPhoneMasked string
	PhoneVerified      bool
}

type AccountReader interface {
	AccountInfo(ctx context.Context, accountID string) (AccountInfo, error)
	MarkProfessional(ctx context.Context, accountID string) error
	RevokeSessions(ctx context.Context, accountID string) error
	RevokeSessionsTx(ctx context.Context, q db.Querier, accountID string) error
}

type CategoryInfo struct {
	ID     string
	Slug   string
	Name   string
	Active bool
}

type CatalogReader interface {
	ActiveCategory(ctx context.Context, categoryID string) (CategoryInfo, error)
	ActiveCategories(ctx context.Context, ids []string) ([]CategoryInfo, error)
}

type Repo interface {
	Create(ctx context.Context, q db.Querier, p domain.Profile) error
	Update(ctx context.Context, q db.Querier, p domain.Profile) error
	GetByAccountID(ctx context.Context, q db.Querier, accountID string) (*domain.Profile, error)
	GetByID(ctx context.Context, q db.Querier, profileID string) (*domain.Profile, error)
	ReplaceCategories(ctx context.Context, q db.Querier, profileID string, categoryIDs []string, at time.Time) error
	List(ctx context.Context, q db.Querier, f ListFilter) (ListPage, error)
	PurgeExpiredLocations(ctx context.Context, q db.Querier, now time.Time) ([]string, error)
}

type ListFilter struct {
	Status string
	Query  string
	Cursor string
	Limit  int
}

type ListPage struct {
	Items      []AdminSummary
	NextCursor string
}

type AdminSummary struct {
	ID           string
	AccountID    string
	DisplayName  string
	Status       string
	City         string
	Neighborhood string
	UpdatedAt    time.Time
}

type Config struct {
	PreciseLocationTTL time.Duration
}

type Service struct {
	uow      *db.UoW
	repo     Repo
	accounts AccountReader
	catalog  CatalogReader
	audit    *audit.Writer
	outbox   *outbox.Writer
	clock    clock.Clock
	metrics  *metrics.Registry
	cipher   *security.Cipher
	cfg      Config
}

func NewService(uow *db.UoW, repo Repo, accounts AccountReader, catalog CatalogReader, auditWriter *audit.Writer, outboxWriter *outbox.Writer, c clock.Clock, m *metrics.Registry, cipher *security.Cipher, cfg Config) *Service {
	if c == nil {
		c = clock.System{}
	}
	if cfg.PreciseLocationTTL <= 0 {
		cfg.PreciseLocationTTL = 90 * 24 * time.Hour
	}
	return &Service{uow: uow, repo: repo, accounts: accounts, catalog: catalog, audit: auditWriter,
		outbox: outboxWriter, clock: c, metrics: m, cipher: cipher, cfg: cfg}
}

type CreateInput struct {
	DisplayName       string
	Bio               string
	ServiceType       string
	CategoryIDs       []string
	City              string
	Neighborhood      string
	RadiusKm          *int
	Availability      []domain.Window
	AvailabilityNote  string
	ContactPreference *string
	PreciseLat        *float64
	PreciseLng        *float64
}

type UpdateInput struct {
	DisplayName          *string
	Bio                  *string
	ServiceType          *string
	CategoryIDs          *[]string
	City                 *string
	Neighborhood         *string
	RadiusKm             *int
	ClearRadius          bool
	Availability         *[]domain.Window
	AvailabilityNote     *string
	ContactPreference    *string
	PreciseLat           *float64
	PreciseLng           *float64
	ClearPreciseLocation bool
}

func (s *Service) CreateDraft(ctx context.Context, accountID string, in CreateInput, correlationID string) (*domain.Profile, error) {
	if _, err := s.repo.GetByAccountID(ctx, s.uow.Read(), accountID); err == nil {
		return nil, contracts.New(contracts.CodeConflict, "Esta conta já possui um perfil profissional.")
	} else if !errors.Is(err, ErrNotFound) {
		return nil, err
	}
	acct, err := s.accounts.AccountInfo(ctx, accountID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	name := strings.TrimSpace(in.DisplayName)
	if name == "" {
		name = acct.DisplayName
	}
	if err := domain.ValidateDisplayName(name); err != nil {
		return nil, contracts.Validation(err.Error())
	}
	serviceType := domain.ServiceType(in.ServiceType)
	if serviceType == "" {
		serviceType = domain.ServiceInPerson
	}
	if !serviceType.Valid() {
		return nil, contracts.Validation("Escolha um tipo de atendimento válido.")
	}
	if len([]rune(in.Bio)) > 500 {
		return nil, contracts.Validation("A bio deve ter no máximo 500 caracteres.")
	}
	if err := domain.ValidateWindows(in.Availability); err != nil {
		return nil, contracts.New(contracts.CodeInvalidAvailability, err.Error())
	}
	categoryIDs, err := s.validateCategories(ctx, in.CategoryIDs)
	if err != nil {
		return nil, err
	}
	var override *domain.ContactPreference
	if in.ContactPreference != nil && strings.TrimSpace(*in.ContactPreference) != "" {
		p := domain.ContactPreference(strings.TrimSpace(*in.ContactPreference))
		if !p.Valid() {
			return nil, contracts.Validation("Preferência de contato inválida.")
		}
		override = &p
	}
	now := s.clock.Now()
	profile := domain.Profile{
		ID: id.New("prf"), AccountID: accountID, DisplayName: name, Bio: strings.TrimSpace(in.Bio),
		ServiceType: serviceType, Status: domain.StatusDraft,
		City: strings.TrimSpace(in.City), Neighborhood: strings.TrimSpace(in.Neighborhood),
		PublicAreaLabel: publicAreaLabel(in.City, in.Neighborhood),
		RadiusKm:        in.RadiusKm, Availability: in.Availability, AvailabilityNote: strings.TrimSpace(in.AvailabilityNote),
		ContactPreferenceOverride: override, Verification: verificationFor(acct),
		CategoryIDs: categoryIDs, CreatedAt: now, UpdatedAt: now,
	}
	if profile.City != "" && profile.Neighborhood != "" {
		profile.PublicAreaLabel = profile.City + " · " + profile.Neighborhood
	}
	if err := validateProfileText(&profile); err != nil {
		return nil, err
	}
	if err := s.applyPreciseLocation(&profile, in.PreciseLat, in.PreciseLng, false, now); err != nil {
		return nil, err
	}
	// Promote before creating the profile so a partial failure is recoverable:
	// retrying the POST still succeeds even if the account role was already set.
	if err := s.accounts.MarkProfessional(ctx, accountID); err != nil {
		return nil, err
	}
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Create(ctx, q, profile); err != nil {
			return err
		}
		if err := s.repo.ReplaceCategories(ctx, q, profile.ID, categoryIDs, now); err != nil {
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: accountID, ActorRole: "professional", Action: "directory.profile_created",
			TargetType: "profile", TargetID: profile.ID, CorrelationID: correlationID, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "directory.profile_created", map[string]any{
			"profile_id": profile.ID, "occurred_at": now,
		})
	})
	if errors.Is(err, ErrProfileExists) {
		return nil, contracts.New(contracts.CodeConflict, "Esta conta já possui um perfil profissional.")
	}
	if err != nil {
		return nil, err
	}
	s.count("profile_created")
	return &profile, nil
}

func (s *Service) UpdateOwner(ctx context.Context, accountID string, in UpdateInput, correlationID string) (*domain.Profile, error) {
	profile, err := s.repo.GetByAccountID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if profile.Status == domain.StatusSuspended {
		return nil, contracts.New(contracts.CodeForbidden, "O perfil está suspenso; fale com a operação para revisar.")
	}
	acct, err := s.accounts.AccountInfo(ctx, accountID)
	if err != nil {
		return nil, err
	}
	changed := []string{}
	if in.DisplayName != nil {
		if err := domain.ValidateDisplayName(*in.DisplayName); err != nil {
			return nil, contracts.Validation(err.Error())
		}
		profile.DisplayName = strings.TrimSpace(*in.DisplayName)
		changed = append(changed, "display_name")
	}
	if in.Bio != nil {
		if len([]rune(*in.Bio)) > 500 {
			return nil, contracts.Validation("A bio deve ter no máximo 500 caracteres.")
		}
		profile.Bio = strings.TrimSpace(*in.Bio)
		changed = append(changed, "bio")
	}
	if in.ServiceType != nil {
		st := domain.ServiceType(strings.TrimSpace(*in.ServiceType))
		if !st.Valid() {
			return nil, contracts.Validation("Escolha um tipo de atendimento válido.")
		}
		profile.ServiceType = st
		changed = append(changed, "service_type")
	}
	if in.City != nil {
		profile.City = strings.TrimSpace(*in.City)
		changed = append(changed, "city")
	}
	if in.Neighborhood != nil {
		profile.Neighborhood = strings.TrimSpace(*in.Neighborhood)
		changed = append(changed, "neighborhood")
	}
	if in.ClearRadius {
		profile.RadiusKm = nil
		changed = append(changed, "radius_km")
	} else if in.RadiusKm != nil {
		if *in.RadiusKm < 1 || *in.RadiusKm > 50 {
			return nil, contracts.New(contracts.CodeInvalidServiceArea, "")
		}
		radius := *in.RadiusKm
		profile.RadiusKm = &radius
		changed = append(changed, "radius_km")
	}
	if in.Availability != nil {
		if err := domain.ValidateWindows(*in.Availability); err != nil {
			return nil, contracts.New(contracts.CodeInvalidAvailability, err.Error())
		}
		profile.Availability = *in.Availability
		changed = append(changed, "availability")
	}
	if in.AvailabilityNote != nil {
		profile.AvailabilityNote = strings.TrimSpace(*in.AvailabilityNote)
		changed = append(changed, "availability_note")
	}
	if in.ContactPreference != nil {
		raw := strings.TrimSpace(*in.ContactPreference)
		if raw == "" {
			profile.ContactPreferenceOverride = nil
		} else {
			p := domain.ContactPreference(raw)
			if !p.Valid() {
				return nil, contracts.Validation("Preferência de contato inválida.")
			}
			profile.ContactPreferenceOverride = &p
		}
		changed = append(changed, "contact_preference")
	}
	if in.CategoryIDs != nil {
		ids, err := s.validateCategories(ctx, *in.CategoryIDs)
		if err != nil {
			return nil, err
		}
		profile.CategoryIDs = ids
		changed = append(changed, "categories")
	}
	profile.PublicAreaLabel = ""
	if profile.City != "" || profile.Neighborhood != "" {
		profile.PublicAreaLabel = publicAreaLabel(profile.City, profile.Neighborhood)
	}
	profile.Verification = verificationFor(acct)
	now := s.clock.Now()
	if err := validateProfileText(profile); err != nil {
		return nil, err
	}
	if err := s.applyPreciseLocation(profile, in.PreciseLat, in.PreciseLng, in.ClearPreciseLocation, now); err != nil {
		return nil, err
	}
	if in.PreciseLat != nil || in.ClearPreciseLocation {
		changed = append(changed, "precise_location")
	}
	profile.UpdatedAt = now
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Update(ctx, q, *profile); err != nil {
			return err
		}
		if in.CategoryIDs != nil {
			if err := s.repo.ReplaceCategories(ctx, q, profile.ID, profile.CategoryIDs, now); err != nil {
				return err
			}
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: accountID, ActorRole: "professional", Action: "directory.profile_updated",
			TargetType: "profile", TargetID: profile.ID, CorrelationID: correlationID,
			Metadata: map[string]any{"changed": changed}, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "directory.profile_updated", map[string]any{
			"profile_id": profile.ID, "changed": changed, "occurred_at": now,
		})
	})
	if err != nil {
		return nil, mapWriteError(err)
	}
	s.count("profile_updated")
	return profile, nil
}

func (s *Service) Publish(ctx context.Context, accountID, correlationID string) (*domain.Profile, error) {
	profile, err := s.repo.GetByAccountID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if profile.Status == domain.StatusPublished {
		return profile, nil
	}
	if profile.Status == domain.StatusSuspended {
		return nil, contracts.New(contracts.CodeForbidden, "O perfil está suspenso.")
	}
	if profile.Status == domain.StatusPaused {
		return nil, contracts.New(contracts.CodeInvalidTransition, "Use a reativação para retomar um perfil pausado.")
	}
	acct, err := s.accounts.AccountInfo(ctx, accountID)
	if err != nil {
		return nil, err
	}
	if !acct.EmailVerified {
		return nil, contracts.New(contracts.CodeEmailNotConfirmed, "")
	}
	categories, err := s.activeCategories(ctx, profile.CategoryIDs)
	if err != nil {
		return nil, err
	}
	if len(categories) != len(profile.CategoryIDs) {
		return nil, contracts.New(contracts.CodeProfileIncomplete, "Escolha ao menos uma categoria ativa.")
	}
	if err := validateProfileText(profile); err != nil {
		return nil, err
	}
	if err := s.validatePublishable(profile); err != nil {
		return nil, err
	}
	now := s.clock.Now()
	profile.Status = domain.StatusPublished
	if profile.PublishedAt == nil {
		profile.PublishedAt = &now
	}
	profile.Verification = domain.VerificationContactVerified
	profile.UpdatedAt = now
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Update(ctx, q, *profile); err != nil {
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: accountID, ActorRole: "professional", Action: "directory.profile_published",
			TargetType: "profile", TargetID: profile.ID, CorrelationID: correlationID, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "directory.profile_published", map[string]any{
			"profile_id": profile.ID, "occurred_at": now,
		})
	})
	if err != nil {
		return nil, mapWriteError(err)
	}
	s.count("profile_published")
	return profile, nil
}

func (s *Service) Pause(ctx context.Context, accountID, correlationID string) (*domain.Profile, error) {
	profile, err := s.repo.GetByAccountID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if profile.Status != domain.StatusPublished {
		return nil, contracts.New(contracts.CodeInvalidTransition, "")
	}
	now := s.clock.Now()
	profile.Status = domain.StatusPaused
	profile.UpdatedAt = now
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Update(ctx, q, *profile); err != nil {
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: accountID, ActorRole: "professional", Action: "directory.profile_paused",
			TargetType: "profile", TargetID: profile.ID, CorrelationID: correlationID, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "directory.profile_paused", map[string]any{
			"profile_id": profile.ID, "occurred_at": now,
		})
	})
	if err != nil {
		return nil, mapWriteError(err)
	}
	s.count("profile_paused")
	return profile, nil
}

func (s *Service) Resume(ctx context.Context, accountID, correlationID string) (*domain.Profile, error) {
	profile, err := s.repo.GetByAccountID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if profile.Status != domain.StatusPaused {
		return nil, contracts.New(contracts.CodeInvalidTransition, "")
	}
	categories, err := s.activeCategories(ctx, profile.CategoryIDs)
	if err != nil {
		return nil, err
	}
	if len(categories) != len(profile.CategoryIDs) {
		return nil, contracts.New(contracts.CodeProfileIncomplete, "Escolha ao menos uma categoria ativa.")
	}
	if err := validateProfileText(profile); err != nil {
		return nil, err
	}
	if err := s.validatePublishable(profile); err != nil {
		return nil, err
	}
	now := s.clock.Now()
	profile.Status = domain.StatusPublished
	profile.UpdatedAt = now
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Update(ctx, q, *profile); err != nil {
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: accountID, ActorRole: "professional", Action: "directory.profile_resumed",
			TargetType: "profile", TargetID: profile.ID, CorrelationID: correlationID, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "directory.profile_resumed", map[string]any{
			"profile_id": profile.ID, "occurred_at": now,
		})
	})
	if err != nil {
		return nil, mapWriteError(err)
	}
	s.count("profile_resumed")
	return profile, nil
}

func (s *Service) validatePublishable(profile *domain.Profile) error {
	if err := profile.ValidateForPublish(); err != nil {
		msg := err.Error()
		switch {
		case strings.Contains(msg, "categoria"), strings.Contains(msg, "nome"), strings.Contains(msg, "tipo"):
			return contracts.New(contracts.CodeProfileIncomplete, msg)
		case strings.Contains(msg, "cidade"), strings.Contains(msg, "bairro"), strings.Contains(msg, "raio"):
			return contracts.New(contracts.CodeInvalidServiceArea, msg)
		case strings.Contains(msg, "janela"), strings.Contains(msg, "fuso"), strings.Contains(msg, "hora"):
			return contracts.New(contracts.CodeInvalidAvailability, msg)
		default:
			return contracts.New(contracts.CodeProfileIncomplete, msg)
		}
	}
	return nil
}

func (s *Service) validateCategories(ctx context.Context, ids []string) ([]string, error) {
	out := make([]string, 0, len(ids))
	seen := map[string]bool{}
	for _, categoryID := range ids {
		categoryID = strings.TrimSpace(categoryID)
		if categoryID == "" || seen[categoryID] {
			continue
		}
		if _, err := s.catalog.ActiveCategory(ctx, categoryID); err != nil {
			return nil, contracts.Validation("Escolha uma categoria ativa.")
		}
		seen[categoryID] = true
		out = append(out, categoryID)
	}
	return out, nil
}

func (s *Service) activeCategories(ctx context.Context, ids []string) ([]CategoryInfo, error) {
	out := make([]CategoryInfo, 0, len(ids))
	for _, categoryID := range ids {
		info, err := s.catalog.ActiveCategory(ctx, categoryID)
		if err != nil {
			continue
		}
		out = append(out, info)
	}
	return out, nil
}

func (s *Service) applyPreciseLocation(p *domain.Profile, lat, lng *float64, clear bool, now time.Time) error {
	if clear {
		p.PreciseLatEnc, p.PreciseLngEnc = "", ""
		p.PreciseLocationExpiresAt = nil
		return nil
	}
	if lat == nil && lng == nil {
		return nil
	}
	if lat == nil || lng == nil {
		return contracts.Validation("Informe latitude e longitude juntas.")
	}
	if !domain.ValidLatLng(*lat, *lng) {
		return contracts.Validation("Coordenada inválida.")
	}
	if s.cipher == nil {
		return contracts.New(contracts.CodeInternal, "Criptografia de localização indisponível.")
	}
	encLat, err := s.cipher.EncryptString(fmt.Sprintf("%.6f", *lat))
	if err != nil {
		return err
	}
	encLng, err := s.cipher.EncryptString(fmt.Sprintf("%.6f", *lng))
	if err != nil {
		return err
	}
	expires := now.Add(s.cfg.PreciseLocationTTL)
	p.PreciseLatEnc, p.PreciseLngEnc = encLat, encLng
	p.PreciseLocationExpiresAt = &expires
	return nil
}

type PublicView struct {
	ID               string
	DisplayName      string
	Bio              string
	ServiceType      string
	Categories       []CategoryView
	City             string
	Neighborhood     string
	RadiusKm         *int
	Availability     []domain.Window
	AvailabilityNote string
	Verification     string
	PublishedAt      *time.Time
}

type CategoryView struct {
	ID   string
	Slug string
	Name string
}

type OwnerView struct {
	PublicView
	Status                   string
	CategoryIDs              []string
	Contact                  domain.ContactView
	PreciseLocationSet       bool
	PreciseLocationExpiresAt *time.Time
	SuspendedReason          string
}

func (s *Service) PublicView(ctx context.Context, profileID string) (*PublicView, error) {
	profile, err := s.repo.GetByID(ctx, s.uow.Read(), profileID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if profile.Status != domain.StatusPublished {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if err := validateProfileText(profile); err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if err := s.validatePublishable(profile); err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	categories, err := s.activeCategories(ctx, profile.CategoryIDs)
	if err != nil || len(categories) != len(profile.CategoryIDs) {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	return publicView(profile, categories), nil
}

func (s *Service) OwnerView(ctx context.Context, accountID string) (*OwnerView, error) {
	profile, err := s.repo.GetByAccountID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	acct, err := s.accounts.AccountInfo(ctx, accountID)
	if err != nil {
		return nil, err
	}
	categories := make([]CategoryInfo, 0, len(profile.CategoryIDs))
	for _, categoryID := range profile.CategoryIDs {
		if info, err := s.catalog.ActiveCategory(ctx, categoryID); err == nil {
			categories = append(categories, info)
		}
	}
	view := publicView(profile, categories)
	override := ""
	if profile.ContactPreferenceOverride != nil {
		override = string(*profile.ContactPreferenceOverride)
	}
	contact := domain.ResolveContact(acct.ContactPreference, override, acct.ContactEmailMasked, acct.ContactPhoneMasked, acct.PhoneVerified)
	return &OwnerView{
		PublicView: *view, Status: string(profile.Status), CategoryIDs: profile.CategoryIDs,
		Contact: contact, PreciseLocationSet: profile.HasPreciseLocation(),
		PreciseLocationExpiresAt: profile.PreciseLocationExpiresAt, SuspendedReason: profile.SuspendedReason,
	}, nil
}

func (s *Service) Preview(ctx context.Context, accountID string) (*PublicView, error) {
	profile, err := s.repo.GetByAccountID(ctx, s.uow.Read(), accountID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if err := validateProfileText(profile); err != nil {
		return nil, err
	}
	categories := make([]CategoryInfo, 0, len(profile.CategoryIDs))
	for _, categoryID := range profile.CategoryIDs {
		if info, err := s.catalog.ActiveCategory(ctx, categoryID); err == nil {
			categories = append(categories, info)
		}
	}
	return publicView(profile, categories), nil
}

// PreviewByID enforces ownership before returning the same public projection.
func (s *Service) PreviewByID(ctx context.Context, accountID, profileID string) (*PublicView, error) {
	profile, err := s.repo.GetByID(ctx, s.uow.Read(), profileID)
	if err != nil || profile.AccountID != accountID {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if err := validateProfileText(profile); err != nil {
		return nil, err
	}
	categories := make([]CategoryInfo, 0, len(profile.CategoryIDs))
	for _, categoryID := range profile.CategoryIDs {
		if info, err := s.catalog.ActiveCategory(ctx, categoryID); err == nil {
			categories = append(categories, info)
		}
	}
	return publicView(profile, categories), nil
}

func (s *Service) Suspend(ctx context.Context, actorID, actorRole, profileID, reason, evidence, correlationID string) error {
	reason = strings.TrimSpace(reason)
	if len([]rune(reason)) < 3 || len([]rune(reason)) > 500 {
		return contracts.Validation("Informe um motivo entre 3 e 500 caracteres.")
	}
	profile, err := s.repo.GetByID(ctx, s.uow.Read(), profileID)
	if err != nil {
		return contracts.New(contracts.CodeNotFound, "")
	}
	now := s.clock.Now()
	profile.Status = domain.StatusSuspended
	profile.SuspendedAt = &now
	profile.SuspendedReason = reason
	profile.UpdatedAt = now
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Update(ctx, q, *profile); err != nil {
			return err
		}
		// Session revocation is part of the same transaction as suspension, so a
		// revoked profile cannot keep an active session due to a partial failure.
		if err := s.accounts.RevokeSessionsTx(ctx, q, profile.AccountID); err != nil {
			return err
		}
		metadata := map[string]any{}
		if evidence != "" {
			metadata["evidence"] = evidence
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: actorID, ActorRole: actorRole, Action: "directory.profile_suspended",
			TargetType: "profile", TargetID: profile.ID, Reason: reason,
			CorrelationID: correlationID, Metadata: metadata, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "directory.profile_suspended", map[string]any{
			"profile_id": profile.ID, "occurred_at": now,
		})
	})
	return mapWriteError(err)
}

func (s *Service) SuspendAndRevoke(ctx context.Context, actorID, actorRole, profileID, reason, evidence, correlationID string) error {
	return s.Suspend(ctx, actorID, actorRole, profileID, reason, evidence, correlationID)
}

func (s *Service) Reactivate(ctx context.Context, actorID, actorRole, profileID, correlationID string) error {
	profile, err := s.repo.GetByID(ctx, s.uow.Read(), profileID)
	if err != nil {
		return contracts.New(contracts.CodeNotFound, "")
	}
	if profile.Status != domain.StatusSuspended {
		return contracts.New(contracts.CodeInvalidTransition, "")
	}
	categories, err := s.activeCategories(ctx, profile.CategoryIDs)
	if err != nil || len(categories) != len(profile.CategoryIDs) {
		return contracts.New(contracts.CodeProfileIncomplete, "Ajuste as categorias antes de reativar.")
	}
	if err := validateProfileText(profile); err != nil {
		return err
	}
	if err := s.validatePublishable(profile); err != nil {
		return err
	}
	now := s.clock.Now()
	profile.Status = domain.StatusPublished
	profile.SuspendedAt = nil
	profile.SuspendedReason = ""
	profile.UpdatedAt = now
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Update(ctx, q, *profile); err != nil {
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: actorID, ActorRole: actorRole, Action: "directory.profile_reactivated",
			TargetType: "profile", TargetID: profile.ID, CorrelationID: correlationID, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "directory.profile_reactivated", map[string]any{
			"profile_id": profile.ID, "occurred_at": now,
		})
	})
	return mapWriteError(err)
}

func (s *Service) ListAdmin(ctx context.Context, f ListFilter) (ListPage, error) {
	return s.repo.List(ctx, s.uow.Read(), f)
}

// PurgeExpiredLocations removes precise coordinates past retention and records
// an auditable event per profile.
func (s *Service) PurgeExpiredLocations(ctx context.Context, now time.Time) (int, error) {
	var ids []string
	err := s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		var err error
		ids, err = s.repo.PurgeExpiredLocations(ctx, q, now)
		if err != nil {
			return err
		}
		for _, profileID := range ids {
			if err := s.audit.Write(ctx, q, audit.Event{
				ActorID: "system", ActorRole: "system", Action: "directory.location_purged",
				TargetType: "profile", TargetID: profileID, Reason: "retenção de 90 dias",
				CreatedAt: now,
			}); err != nil {
				return err
			}
		}
		if len(ids) > 0 {
			if err := s.outbox.Publish(ctx, q, "directory.locations_purged", map[string]any{
				"count": len(ids), "occurred_at": now,
			}); err != nil {
				return err
			}
		}
		return nil
	})
	if err == nil && len(ids) > 0 {
		s.count("location_purged")
	}
	return len(ids), err
}

func publicView(p *domain.Profile, categories []CategoryInfo) *PublicView {
	views := make([]CategoryView, 0, len(categories))
	for _, c := range categories {
		views = append(views, CategoryView{ID: c.ID, Slug: c.Slug, Name: c.Name})
	}
	return &PublicView{
		ID: p.ID, DisplayName: p.DisplayName, Bio: p.Bio, ServiceType: string(p.ServiceType),
		Categories: views, City: p.City, Neighborhood: p.Neighborhood, RadiusKm: p.RadiusKm,
		Availability: p.Availability, AvailabilityNote: p.AvailabilityNote,
		Verification: string(p.Verification), PublishedAt: p.PublishedAt,
	}
}

func verificationFor(acct AccountInfo) domain.Verification {
	if acct.EmailVerified {
		return domain.VerificationContactVerified
	}
	return domain.VerificationUnverified
}

func publicAreaLabel(city, neighborhood string) string {
	city, neighborhood = strings.TrimSpace(city), strings.TrimSpace(neighborhood)
	switch {
	case city != "" && neighborhood != "":
		return city + " · " + neighborhood
	case city != "":
		return city
	default:
		return neighborhood
	}
}

func (s *Service) count(name string) {
	if s.metrics != nil {
		s.metrics.Inc(name)
	}
}

func mapWriteError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, ErrConflict):
		return contracts.New(contracts.CodeConflict, "O perfil foi alterado por outra sessão. Recarregue e tente novamente.")
	case errors.Is(err, ErrProfileExists):
		return contracts.New(contracts.CodeConflict, "Esta conta já possui um perfil profissional.")
	default:
		return err
	}
}

func validateProfileText(p *domain.Profile) error {
	fields := []struct {
		name  string
		value string
		max   int
	}{
		{"nome de exibição", p.DisplayName, 80},
		{"bio", p.Bio, 500},
		{"cidade", p.City, 80},
		{"bairro", p.Neighborhood, 80},
		{"disponibilidade", p.AvailabilityNote, 200},
	}
	for _, f := range fields {
		if err := domain.ValidatePublicText(f.name, f.value, f.max); err != nil {
			var textErr *domain.PublicTextError
			if errors.As(err, &textErr) && textErr.Reason == "pii" {
				return contracts.New(contracts.CodePIINotAllowed, textErr.Error())
			}
			return contracts.Validation(err.Error())
		}
	}
	return nil
}
