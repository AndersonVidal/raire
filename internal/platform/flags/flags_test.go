package flags

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/raire/raire/internal/platform/db"
)

func TestFlags(t *testing.T) {
	database, err := db.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	ctx := context.Background()
	if err := db.Migrate(ctx, database); err != nil {
		t.Fatal(err)
	}
	if err := db.Seed(ctx, database); err != nil {
		t.Fatal(err)
	}
	r := &Repo{DB: database}
	if r.IsEnabled(ctx, "unknown.flag") {
		t.Fatal("unknown flag enabled")
	}
	rows, err := r.List(ctx, database)
	if err != nil || len(rows) == 0 {
		t.Fatalf("list rows=%d err=%v", len(rows), err)
	}
	if err := r.Update(ctx, database, "catalog.use_synonyms", false, time.Now().UTC()); err != nil {
		t.Fatal(err)
	}
	if r.IsEnabled(ctx, "catalog.use_synonyms") {
		t.Fatal("flag did not change")
	}
}
