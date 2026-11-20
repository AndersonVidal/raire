package idempotency

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/raire/raire/internal/platform/db"
)

func TestStore(t *testing.T) {
	database, err := db.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	ctx := context.Background()
	if err := db.Migrate(ctx, database); err != nil {
		t.Fatal(err)
	}
	store := &Store{DB: database}
	now := time.Now().UTC()
	if err := store.Check(ctx, database, "", "op", "hash", "", now); err != nil {
		t.Fatal(err)
	}
	if err := store.Check(ctx, database, "key", "op", "hash", "", now); err != nil {
		t.Fatal(err)
	}
	if err := store.Check(ctx, database, "key", "op", "hash", "", now); !errors.Is(err, ErrReplay) {
		t.Fatalf("expected replay, got %v", err)
	}
	if err := store.Check(ctx, database, "key", "op", "other", "", now); !errors.Is(err, ErrConflict) {
		t.Fatalf("expected conflict, got %v", err)
	}
}
