package audit

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/raire/raire/internal/platform/clock"
	"github.com/raire/raire/internal/platform/db"
	"github.com/raire/raire/internal/platform/id"
)

// Event is an append-only administrative or security-relevant fact.
type Event struct {
	ID            string
	ActorID       string
	ActorRole     string
	Action        string
	TargetType    string
	TargetID      string
	Reason        string
	CorrelationID string
	Metadata      map[string]any
	CreatedAt     time.Time
}

type Writer struct {
	clock clock.Clock
}

func NewWriter(c clock.Clock) *Writer {
	if c == nil {
		c = clock.System{}
	}
	return &Writer{clock: c}
}

func (w *Writer) Write(ctx context.Context, q db.Querier, e Event) error {
	if e.Action == "" || e.TargetType == "" || e.TargetID == "" {
		return fmt.Errorf("audit event requires action and target")
	}
	if e.ID == "" {
		e.ID = id.New("aud")
	}
	if e.CreatedAt.IsZero() {
		e.CreatedAt = w.clock.Now()
	}
	metadata := e.Metadata
	if metadata == nil {
		metadata = map[string]any{}
	}
	raw, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	_, err = q.ExecContext(ctx, `INSERT INTO audit_events
		(id, actor_id, actor_role, action, target_type, target_id, reason, correlation_id, metadata_json, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		e.ID, nullString(e.ActorID), e.ActorRole, e.Action, e.TargetType, e.TargetID,
		e.Reason, e.CorrelationID, string(raw), e.CreatedAt.UTC().Format(time.RFC3339Nano))
	return err
}

type Filter struct {
	ActorID    string
	TargetType string
	TargetID   string
	From       *time.Time
	To         *time.Time
	Cursor     string
	Limit      int
}

type Page struct {
	Items      []Event
	NextCursor string
}

// List returns audit events with cursor pagination. Audit is append-only; there
// is intentionally no update or delete API.
func (w *Writer) List(ctx context.Context, q db.Querier, f Filter) (Page, error) {
	limit := f.Limit
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	where := []string{"1=1"}
	args := []any{}
	if f.ActorID != "" {
		where = append(where, "actor_id = ?")
		args = append(args, f.ActorID)
	}
	if f.TargetType != "" {
		where = append(where, "target_type = ?")
		args = append(args, f.TargetType)
	}
	if f.TargetID != "" {
		where = append(where, "target_id = ?")
		args = append(args, f.TargetID)
	}
	if f.From != nil {
		where = append(where, "created_at >= ?")
		args = append(args, f.From.UTC().Format(time.RFC3339Nano))
	}
	if f.To != nil {
		where = append(where, "created_at <= ?")
		args = append(args, f.To.UTC().Format(time.RFC3339Nano))
	}
	if f.Cursor != "" {
		at, id, err := decodeCursor(f.Cursor)
		if err != nil {
			return Page{}, err
		}
		where = append(where, "(created_at < ? OR (created_at = ? AND id < ?))")
		args = append(args, at, at, id)
	}
	args = append(args, limit+1)
	query := `SELECT id, coalesce(actor_id,''), actor_role, action, target_type, target_id,
		reason, correlation_id, metadata_json, created_at
		FROM audit_events WHERE ` + strings.Join(where, " AND ") + `
		ORDER BY created_at DESC, id DESC LIMIT ?`
	rows, err := q.QueryContext(ctx, query, args...)
	if err != nil {
		return Page{}, err
	}
	defer rows.Close()

	page := Page{}
	for rows.Next() {
		var e Event
		var actor sql.NullString
		var metadata string
		var created string
		if err := rows.Scan(&e.ID, &actor, &e.ActorRole, &e.Action, &e.TargetType, &e.TargetID,
			&e.Reason, &e.CorrelationID, &metadata, &created); err != nil {
			return Page{}, err
		}
		e.ActorID = actor.String
		_ = json.Unmarshal([]byte(metadata), &e.Metadata)
		e.CreatedAt, _ = time.Parse(time.RFC3339Nano, created)
		page.Items = append(page.Items, e)
	}
	if err := rows.Err(); err != nil {
		return Page{}, err
	}
	if len(page.Items) > limit {
		last := page.Items[limit-1]
		page.NextCursor = encodeCursor(last.CreatedAt, last.ID)
		page.Items = page.Items[:limit]
	}
	return page, nil
}

func encodeCursor(at time.Time, id string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(at.UTC().Format(time.RFC3339Nano) + "|" + id))
}

func decodeCursor(cursor string) (string, string, error) {
	raw, err := base64.RawURLEncoding.DecodeString(cursor)
	if err != nil {
		return "", "", fmt.Errorf("invalid cursor")
	}
	parts := strings.SplitN(string(raw), "|", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid cursor")
	}
	return parts[0], parts[1], nil
}

func nullString(v string) any {
	if v == "" {
		return nil
	}
	return v
}
