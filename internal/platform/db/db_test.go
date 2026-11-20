package db

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/raire/raire/internal/platform/id"
)

func TestMigrateAndSeedIdempotent(t *testing.T) {
	database, err := Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	ctx := context.Background()
	if err := Ready(ctx, database); err == nil {
		t.Fatal("empty database reported ready")
	}
	if err := Migrate(ctx, database); err != nil {
		t.Fatal(err)
	}
	if err := Seed(ctx, database); err != nil {
		t.Fatal(err)
	}
	var first int
	if err := database.QueryRowContext(ctx, `SELECT count(*) FROM categories`).Scan(&first); err != nil {
		t.Fatal(err)
	}
	if err := Seed(ctx, database); err != nil {
		t.Fatal(err)
	}
	var second int
	if err := database.QueryRowContext(ctx, `SELECT count(*) FROM categories`).Scan(&second); err != nil {
		t.Fatal(err)
	}
	if first != second || first < 5 {
		t.Fatalf("seed not idempotent: %d -> %d", first, second)
	}
	if err := Migrate(ctx, database); err != nil {
		t.Fatal(err)
	}
	if err := Ready(ctx, database); err != nil {
		t.Fatal(err)
	}
}

func TestAuditAppendOnly(t *testing.T) {
	database, err := Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	ctx := context.Background()
	if err := Migrate(ctx, database); err != nil {
		t.Fatal(err)
	}
	eventID := id.New("aud")
	if _, err := database.ExecContext(ctx, `INSERT INTO audit_events
		(id, actor_id, actor_role, action, target_type, target_id, created_at)
		VALUES (?, 'acc_1', 'admin', 'test.action', 'account', 'acc_1', '2026-11-20T00:00:00Z')`, eventID); err != nil {
		t.Fatal(err)
	}
	if _, err := database.ExecContext(ctx, `UPDATE audit_events SET reason = 'x' WHERE id = ?`, eventID); err == nil {
		t.Fatal("audit update succeeded")
	}
	if _, err := database.ExecContext(ctx, `DELETE FROM audit_events WHERE id = ?`, eventID); err == nil {
		t.Fatal("audit delete succeeded")
	}
}
