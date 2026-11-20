package app

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/raire/raire/internal/catalog/adapters/sqlite"
	catalogapp "github.com/raire/raire/internal/catalog/app"
	"github.com/raire/raire/internal/contracts"
	directorycatalog "github.com/raire/raire/internal/directory/adapters/catalog"
	directoryidentity "github.com/raire/raire/internal/directory/adapters/identity"
	directorysqlite "github.com/raire/raire/internal/directory/adapters/sqlite"
	directoryapp "github.com/raire/raire/internal/directory/app"
	identitysqlite "github.com/raire/raire/internal/identity/adapters/sqlite"
	identityapp "github.com/raire/raire/internal/identity/app"
	"github.com/raire/raire/internal/identity/domain"
	"github.com/raire/raire/internal/platform/audit"
	"github.com/raire/raire/internal/platform/clock"
	"github.com/raire/raire/internal/platform/config"
	"github.com/raire/raire/internal/platform/db"
	"github.com/raire/raire/internal/platform/flags"
	"github.com/raire/raire/internal/platform/id"
	"github.com/raire/raire/internal/platform/idempotency"
	"github.com/raire/raire/internal/platform/metrics"
	"github.com/raire/raire/internal/platform/outbox"
	"github.com/raire/raire/internal/platform/ratelimit"
	"github.com/raire/raire/internal/platform/security"
)

// App wires the foundation contexts. Only this package knows concrete
// adapters; contexts communicate through the narrow interfaces they declare.
type App struct {
	Cfg          config.Config
	DB           *sql.DB
	UoW          *db.UoW
	Identity     *identityapp.Service
	Catalog      *catalogapp.Service
	Directory    *directoryapp.Service
	IdentityRepo *identitysqlite.Repo
	Flags        *flags.Repo
	Audit        *audit.Writer
	Outbox       *outbox.Writer
	Metrics      *metrics.Registry
	Limiter      *ratelimit.Limiter
	Cipher       *security.Cipher
}

func New(cfg config.Config) (*App, error) {
	database, err := db.Open(cfg.DBPath)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := db.Migrate(ctx, database); err != nil {
		_ = database.Close()
		return nil, err
	}
	if err := db.Seed(ctx, database); err != nil {
		_ = database.Close()
		return nil, err
	}
	cipher, err := security.NewCipher(cfg.EncryptionKey)
	if err != nil {
		_ = database.Close()
		return nil, err
	}
	uow := &db.UoW{DB: database}
	clk := clock.System{}
	auditWriter := audit.NewWriter(clk)
	outboxWriter := outbox.NewWriter(clk)
	idem := &idempotency.Store{DB: database}
	metricRegistry := metrics.New()
	limiter := ratelimit.New(clk)
	flagRepo := &flags.Repo{DB: database}

	identityRepo := identitysqlite.New(database)
	identitySvc := identityapp.NewService(uow, identityRepo, auditWriter, outboxWriter, idem, clk, metricRegistry, identityapp.Config{
		SessionTTL: cfg.SessionTTL, ConfirmTokenTTL: cfg.ConfirmTokenTTL, ResetTokenTTL: cfg.ResetTokenTTL,
		TermsVersion: cfg.TermsVersion, PrivacyVersion: cfg.PrivacyVersion, Argon2: argonParams(cfg.Argon2), Cipher: cipher,
	})

	catalogRepo := sqlite.New(database)
	catalogSvc := catalogapp.NewService(uow, catalogRepo, auditWriter, outboxWriter, clk, metricRegistry)

	accountAdapter := directoryidentity.New(identitySvc)
	catalogAdapter := directorycatalog.New(catalogSvc)
	directoryRepo := directorysqlite.New(database)
	directorySvc := directoryapp.NewService(uow, directoryRepo, accountAdapter, catalogAdapter,
		auditWriter, outboxWriter, clk, metricRegistry, cipher, directoryapp.Config{PreciseLocationTTL: cfg.PreciseLocationTTL})

	return &App{
		Cfg: cfg, DB: database, UoW: uow, Identity: identitySvc, Catalog: catalogSvc,
		Directory: directorySvc, IdentityRepo: identityRepo, Flags: flagRepo,
		Audit: auditWriter, Outbox: outboxWriter, Metrics: metricRegistry, Limiter: limiter, Cipher: cipher,
	}, nil
}

func (a *App) Close() error {
	if a == nil || a.DB == nil {
		return nil
	}
	return a.DB.Close()
}

// Ready reports database connectivity and applied migrations without exposing
// the database handle to HTTP handlers.
func (a *App) Ready(ctx context.Context) error {
	if a == nil || a.DB == nil {
		return sql.ErrConnDone
	}
	return db.Ready(ctx, a.DB)
}

// ListAudit is the application query used by the admin HTTP handler.
func (a *App) ListAudit(ctx context.Context, f audit.Filter) (audit.Page, error) {
	return a.Audit.List(ctx, a.DB, f)
}

// ListFlags is the application query used by the admin HTTP handler.
func (a *App) ListFlags(ctx context.Context) ([]flags.Flag, error) {
	return a.Flags.List(ctx, a.DB)
}

func argonParams(c config.Argon2) security.Params {
	return security.Params{
		Memory: c.Memory, Iterations: c.Iterations, Parallelism: c.Parallelism,
		SaltLength: c.SaltLength, KeyLength: c.KeyLength,
	}
}

// CreateOperator provisions a privileged account for local development and
// tests. It must never be exposed by an HTTP route.
func (a *App) CreateOperator(ctx context.Context, email, password string) (*domain.Account, string, []string, error) {
	email = domain.NormalizeEmail(email)
	if err := domain.ValidateEmail(email); err != nil {
		return nil, "", nil, err
	}
	if err := domain.ValidatePassword(password); err != nil {
		return nil, "", nil, err
	}
	if existing, err := a.IdentityRepo.GetAccountByEmail(ctx, a.UoW.Read(), email); err == nil && existing != nil {
		return nil, "", nil, errors.New("account already exists")
	}
	hash, err := security.HashPassword(password, argonParams(a.Cfg.Argon2))
	if err != nil {
		return nil, "", nil, err
	}
	now := time.Now().UTC()
	acct := domain.Account{
		ID: id.New("acc"), EmailNormalized: email, PasswordHash: hash, DisplayName: "Operação",
		Role: domain.RoleOperator, EmailVerifiedAt: &now, ContactEmail: email,
		ContactPreference: domain.ContactPlatform, CreatedAt: now, UpdatedAt: now,
	}
	err = a.UoW.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := a.IdentityRepo.CreateAccount(ctx, q, acct); err != nil {
			return err
		}
		for _, c := range []struct {
			kind    domain.ConsentKind
			version string
		}{
			{domain.ConsentTerms, a.Cfg.TermsVersion}, {domain.ConsentPrivacy, a.Cfg.PrivacyVersion},
		} {
			if err := a.IdentityRepo.CreateConsent(ctx, q, domain.Consent{
				ID: id.New("cns"), AccountID: acct.ID, Kind: c.kind, Version: c.version, GrantedAt: now,
			}); err != nil {
				return err
			}
		}
		if err := a.IdentityRepo.UpsertVerification(ctx, q, acct.ID, "contact_email", "verified", now); err != nil {
			return err
		}
		return a.Audit.Write(ctx, q, audit.Event{
			ActorID: "system", ActorRole: "system", Action: "account.operator_created",
			TargetType: "account", TargetID: acct.ID, CreatedAt: now,
		})
	})
	if err != nil {
		return nil, "", nil, err
	}
	secret, codes, err := a.Identity.ProvisionMFA(ctx, acct.ID)
	if err != nil {
		return nil, "", nil, err
	}
	return &acct, secret, codes, nil
}

// SetFeatureFlag changes a flag, writes an audit event and a domain event in the
// same transaction.
func (a *App) SetFeatureFlag(ctx context.Context, actorID, actorRole, key string, enabled bool, correlationID string) error {
	now := time.Now().UTC()
	err := a.UoW.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := a.Flags.Update(ctx, q, key, enabled, now); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return contracts.New(contracts.CodeNotFound, "")
			}
			return err
		}
		if err := a.Audit.Write(ctx, q, audit.Event{
			ActorID: actorID, ActorRole: actorRole, Action: "platform.feature_flag_changed",
			TargetType: "feature_flag", TargetID: key, CorrelationID: correlationID,
			Metadata: map[string]any{"enabled": enabled}, CreatedAt: now,
		}); err != nil {
			return err
		}
		return a.Outbox.Publish(ctx, q, "platform.feature_flag_changed", map[string]any{
			"key": key, "enabled": enabled, "occurred_at": now,
		})
	})
	if err == nil {
		a.Metrics.Inc("feature_flag_changed")
	}
	return err
}
