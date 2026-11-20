package sqlite

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/raire/raire/internal/directory/app"
	"github.com/raire/raire/internal/directory/domain"
	"github.com/raire/raire/internal/platform/db"
)

func TestProfileUpdateVersionConflict(t *testing.T) {
	database, err := db.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	ctx := context.Background()
	if err := db.Migrate(ctx, database); err != nil {
		t.Fatal(err)
	}
	now := time.Now().UTC().Format(time.RFC3339Nano)
	if _, err := database.ExecContext(ctx, `INSERT INTO accounts
		(id, email_normalized, password_hash, display_name, role, contact_preference, created_at, updated_at)
		VALUES ('acc_test', 'test@example.test', 'hash', 'Teste', 'professional', 'platform', ?, ?)`, now, now); err != nil {
		t.Fatal(err)
	}
	repo := New(database)
	p := domain.Profile{
		ID: "prf_test", AccountID: "acc_test", DisplayName: "Teste", ServiceType: domain.ServiceRemote,
		Status: domain.StatusDraft, Verification: domain.VerificationUnverified,
		CategoryIDs: []string{}, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(),
	}
	if err := repo.Create(ctx, database, p); err != nil {
		t.Fatal(err)
	}
	loaded, err := repo.GetByID(ctx, database, "prf_test")
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Version != 1 {
		t.Fatalf("version = %d", loaded.Version)
	}
	loaded.DisplayName = "Teste 2"
	if err := repo.Update(ctx, database, *loaded); err != nil {
		t.Fatal(err)
	}
	// Simulate an admin suspension after the owner read the profile.
	if _, err := database.ExecContext(ctx, `UPDATE professional_profiles SET status = 'suspended', version = version + 1 WHERE id = 'prf_test'`); err != nil {
		t.Fatal(err)
	}
	loaded.DisplayName = "Lost update"
	if err := repo.Update(ctx, database, *loaded); !errors.Is(err, app.ErrConflict) {
		t.Fatalf("stale update error = %v", err)
	}
	after, err := repo.GetByID(ctx, database, "prf_test")
	if err != nil {
		t.Fatal(err)
	}
	if after.Status != domain.StatusSuspended {
		t.Fatalf("suspension was overwritten: %s", after.Status)
	}
}
