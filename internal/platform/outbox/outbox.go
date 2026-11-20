package outbox

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/raire/raire/internal/platform/clock"
	"github.com/raire/raire/internal/platform/db"
	"github.com/raire/raire/internal/platform/id"
)

// Writer stores a domain event in the same transaction as its aggregate. A
// worker reads pending rows; the MVP deliberately has no external broker.
type Writer struct {
	clock clock.Clock
}

func NewWriter(c clock.Clock) *Writer {
	if c == nil {
		c = clock.System{}
	}
	return &Writer{clock: c}
}

// Publish records an event. Callers must pass IDs and non-personal fields only.
func (w *Writer) Publish(ctx context.Context, q db.Querier, eventType string, payload map[string]any) error {
	if eventType == "" {
		return fmt.Errorf("outbox event type is required")
	}
	if payload == nil {
		payload = map[string]any{}
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	_, err = q.ExecContext(ctx, `INSERT INTO outbox_events
		(id, event_type, payload_json, occurred_at, attempts) VALUES (?, ?, ?, ?, 0)`,
		id.New("evt"), eventType, string(raw), w.clock.Now().UTC().Format(time.RFC3339Nano))
	return err
}

type Event struct {
	ID          string
	Type        string
	Payload     map[string]any
	OccurredAt  time.Time
	ProcessedAt *time.Time
	Attempts    int
	LastError   string
}

// Pending returns unprocessed events in order.
func Pending(ctx context.Context, q db.Querier, limit int) ([]Event, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	rows, err := q.QueryContext(ctx, `SELECT id, event_type, payload_json, occurred_at, processed_at, attempts, coalesce(last_error,'')
		FROM outbox_events WHERE processed_at IS NULL ORDER BY occurred_at, id LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var e Event
		var payload, occurred string
		var processed sql.NullString
		if err := rows.Scan(&e.ID, &e.Type, &payload, &occurred, &processed, &e.Attempts, &e.LastError); err != nil {
			return nil, err
		}
		_ = json.Unmarshal([]byte(payload), &e.Payload)
		e.OccurredAt, _ = time.Parse(time.RFC3339Nano, occurred)
		if processed.Valid {
			t, _ := time.Parse(time.RFC3339Nano, processed.String)
			e.ProcessedAt = &t
		}
		events = append(events, e)
	}
	return events, rows.Err()
}

// MarkProcessed is called by the internal worker after a handler succeeds.
func MarkProcessed(ctx context.Context, q db.Querier, id string, at time.Time) error {
	_, err := q.ExecContext(ctx, `UPDATE outbox_events SET processed_at = ?, attempts = attempts + 1 WHERE id = ?`,
		at.UTC().Format(time.RFC3339Nano), id)
	return err
}
