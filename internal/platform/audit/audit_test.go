package audit

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/raire/raire/internal/platform/clock"
	"github.com/raire/raire/internal/platform/db"
)

func TestWriteAndList(t *testing.T) {
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
	for i := 0; i < 3; i++ {
		if err := w.Write(ctx, database, Event{
			ActorID: "acc_1", ActorRole: "admin", Action: "test.action", TargetType: "profile",
			TargetID: "prf_1", Reason: "teste", CorrelationID: "corr-1",
			Metadata: map[string]any{"n": i}, CreatedAt: now.Add(time.Duration(i) * time.Second),
		}); err != nil {
			t.Fatal(err)
		}
	}
	page, err := w.List(ctx, database, Filter{Limit: 2})
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Items) != 2 || page.NextCursor == "" {
		t.Fatalf("page = %+v", page)
	}
	next, err := w.List(ctx, database, Filter{Limit: 2, Cursor: page.NextCursor})
	if err != nil {
		t.Fatal(err)
	}
	if len(next.Items) != 1 {
		t.Fatalf("next items = %d", len(next.Items))
	}
	if _, err := database.ExecContext(ctx, `UPDATE audit_events SET reason = 'x'`); err == nil {
		t.Fatal("audit update allowed")
	}
}
