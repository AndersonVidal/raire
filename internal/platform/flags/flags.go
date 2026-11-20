package flags

import (
	"context"
	"database/sql"
	"time"

	"github.com/raire/raire/internal/platform/db"
)

type Flag struct {
	Key         string
	Enabled     bool
	Description string
	Owner       string
	ReviewAt    *time.Time
	UpdatedAt   time.Time
}

type Repo struct {
	DB *sql.DB
}

// IsEnabled resolves an unknown flag to false: the safe default is always off.
func (r *Repo) IsEnabled(ctx context.Context, key string) bool {
	var enabled int
	err := r.DB.QueryRowContext(ctx, `SELECT enabled FROM feature_flags WHERE key = ?`, key).Scan(&enabled)
	if err != nil {
		return false
	}
	return enabled == 1
}

func (r *Repo) List(ctx context.Context, q db.Querier) ([]Flag, error) {
	rows, err := q.QueryContext(ctx, `SELECT key, enabled, description, owner, review_at, updated_at
		FROM feature_flags ORDER BY key`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Flag
	for rows.Next() {
		var f Flag
		var enabled int
		var review, updated sql.NullString
		if err := rows.Scan(&f.Key, &enabled, &f.Description, &f.Owner, &review, &updated); err != nil {
			return nil, err
		}
		f.Enabled = enabled == 1
		if review.Valid {
			t, _ := time.Parse(time.RFC3339Nano, review.String)
			f.ReviewAt = &t
		}
		f.UpdatedAt, _ = time.Parse(time.RFC3339Nano, updated.String)
		out = append(out, f)
	}
	return out, rows.Err()
}

func (r *Repo) Update(ctx context.Context, q db.Querier, key string, enabled bool, at time.Time) error {
	res, err := q.ExecContext(ctx, `UPDATE feature_flags SET enabled = ?, updated_at = ? WHERE key = ?`,
		boolInt(enabled), at.UTC().Format(time.RFC3339Nano), key)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func boolInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
