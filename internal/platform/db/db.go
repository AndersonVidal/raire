package db

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"time"

	rairedb "github.com/raire/raire/db"
	_ "modernc.org/sqlite"
)

// Querier is satisfied by *sql.DB and *sql.Tx. Repositories accept it so a
// mutation, its audit event and its outbox event can share one transaction.
type Querier interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// Open opens the SQLite database used locally and in CI.
func Open(path string) (*sql.DB, error) {
	var dsn string
	if path == ":memory:" || path == "" {
		dsn = "file:raire-memory?mode=memory&cache=shared&_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)"
	} else {
		slash := filepath.ToSlash(path)
		switch {
		case strings.HasPrefix(slash, "/"):
			dsn = "file://" + slash
		case filepath.IsAbs(path):
			dsn = "file:///" + slash
		default:
			dsn = "file:" + slash
		}
		dsn += "?_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"
	}
	database, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	if path == ":memory:" || path == "" {
		database.SetMaxOpenConns(1)
	} else {
		database.SetMaxOpenConns(4)
	}
	if err := database.Ping(); err != nil {
		_ = database.Close()
		return nil, err
	}
	return database, nil
}

// Migrate applies pending forward-only migrations inside transactions.
func Migrate(ctx context.Context, database *sql.DB) error {
	if _, err := database.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (
		version TEXT PRIMARY KEY,
		applied_at TEXT NOT NULL
	)`); err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}
	entries, err := fs.ReadDir(rairedb.Migrations, "migrations")
	if err != nil {
		return err
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		var exists int
		if err := database.QueryRowContext(ctx, `SELECT count(*) FROM schema_migrations WHERE version = ?`, e.Name()).Scan(&exists); err != nil {
			return err
		}
		if exists > 0 {
			continue
		}
		body, err := fs.ReadFile(rairedb.Migrations, "migrations/"+e.Name())
		if err != nil {
			return err
		}
		tx, err := database.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, string(body)); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("apply %s: %w", e.Name(), err)
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO schema_migrations (version, applied_at) VALUES (?, ?)`,
			e.Name(), time.Now().UTC().Format(time.RFC3339Nano)); err != nil {
			_ = tx.Rollback()
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

// Seed applies the versioned, idempotent seed files. It is safe to call on
// every boot.
func Seed(ctx context.Context, database *sql.DB) error {
	entries, err := fs.ReadDir(rairedb.Seeds, "seeds")
	if err != nil {
		return err
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		body, err := fs.ReadFile(rairedb.Seeds, "seeds/"+e.Name())
		if err != nil {
			return err
		}
		if _, err := database.ExecContext(ctx, string(body)); err != nil {
			return fmt.Errorf("seed %s: %w", e.Name(), err)
		}
	}
	return nil
}

// Ready verifies connectivity and that migrations were applied.
func Ready(ctx context.Context, database *sql.DB) error {
	if err := database.PingContext(ctx); err != nil {
		return err
	}
	var n int
	if err := database.QueryRowContext(ctx, `SELECT count(*) FROM schema_migrations`).Scan(&n); err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("database has no applied migrations")
	}
	return nil
}

// UoW runs a unit of work in a single transaction.
type UoW struct {
	DB *sql.DB
}

func (u *UoW) Do(ctx context.Context, fn func(context.Context, Querier) error) error {
	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(ctx, tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u *UoW) Read() Querier { return u.DB }
