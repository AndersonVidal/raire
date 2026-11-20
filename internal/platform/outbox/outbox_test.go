package outbox

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/raire/raire/internal/platform/clock"
	"github.com/raire/raire/internal/platform/db"
)

func TestPublishAndProcess(t *testing.T) {
	database, err := db.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	ctx := context.Background()
	if err := db.Migrate(ctx, database); err != nil {
		t.Fatal(err)
	}
	now := time.Unix(1_700_000_000, 0).UTC()
	w := NewWriter(clock.Fixed{T: now})
	if err := w.Publish(ctx, database, "directory.profile_published", map[string]any{"profile_id": "prf_1"}); err != nil {
		t.Fatal(err)
	}
	events, err := Pending(ctx, database, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 1 || events[0].Type != "directory.profile_published" {
		t.Fatalf("events = %+v", events)
	}
	if err := MarkProcessed(ctx, database, events[0].ID, now.Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	events, err = Pending(ctx, database, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 0 {
		t.Fatalf("processed event still pending: %+v", events)
	}
}
