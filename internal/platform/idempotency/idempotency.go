package idempotency

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/raire/raire/internal/platform/db"
)

var (
	ErrReplay   = errors.New("idempotent replay")
	ErrConflict = errors.New("idempotency key reused with a different request")
)

type Store struct {
	DB *sql.DB
}

// Check records a key on first use. A replay with the same request hash returns
// ErrReplay; a different payload returns ErrConflict. No key is a no-op.
func (s *Store) Check(ctx context.Context, q db.Querier, key, operation, requestHash, accountID string, now time.Time) error {
	if key == "" {
		return nil
	}
	if len(key) > 128 {
		return ErrConflict
	}
	var storedHash string
	var expires string
	err := q.QueryRowContext(ctx, `SELECT request_hash, expires_at FROM idempotency_keys WHERE key = ? AND operation = ?`,
		key, operation).Scan(&storedHash, &expires)
	if err == nil {
		exp, _ := time.Parse(time.RFC3339Nano, expires)
		if now.Before(exp) {
			if storedHash == requestHash {
				return ErrReplay
			}
			return ErrConflict
		}
		if _, err := q.ExecContext(ctx, `DELETE FROM idempotency_keys WHERE key = ? AND operation = ?`, key, operation); err != nil {
			return err
		}
	} else if !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	_, err = q.ExecContext(ctx, `INSERT INTO idempotency_keys (key, operation, request_hash, account_id, created_at, expires_at)
		VALUES (?, ?, ?, ?, ?, ?)`, key, operation, requestHash, nullString(accountID),
		now.UTC().Format(time.RFC3339Nano), now.Add(24*time.Hour).UTC().Format(time.RFC3339Nano))
	return err
}

func nullString(v string) any {
	if v == "" {
		return nil
	}
	return v
}
