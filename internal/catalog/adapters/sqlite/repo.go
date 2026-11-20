package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/raire/raire/internal/catalog/app"
	"github.com/raire/raire/internal/catalog/domain"
	"github.com/raire/raire/internal/platform/db"
	"github.com/raire/raire/internal/platform/id"
)

type Repo struct {
	DB *sql.DB
}

func New(database *sql.DB) *Repo { return &Repo{DB: database} }

func (r *Repo) ListActive(ctx context.Context, q db.Querier) ([]domain.Category, error) {
	return r.list(ctx, q, `SELECT id, slug, name, description, active, sort_order, created_at, updated_at
		FROM categories WHERE active = 1 ORDER BY sort_order, name`)
}

func (r *Repo) ListAll(ctx context.Context, q db.Querier) ([]domain.Category, error) {
	return r.list(ctx, q, `SELECT id, slug, name, description, active, sort_order, created_at, updated_at
		FROM categories ORDER BY sort_order, name`)
}

func (r *Repo) get(ctx context.Context, q db.Querier, query string, arg any) (*domain.Category, error) {
	var c domain.Category
	var active int
	var created, updated string
	err := q.QueryRowContext(ctx, query, arg).Scan(&c.ID, &c.Slug, &c.Name, &c.Description, &active,
		&c.SortOrder, &created, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, app.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	c.Active = active == 1
	c.CreatedAt, _ = time.Parse(time.RFC3339Nano, created)
	c.UpdatedAt, _ = time.Parse(time.RFC3339Nano, updated)
	c.Synonyms, err = r.synonyms(ctx, q, c.ID)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *Repo) GetByID(ctx context.Context, q db.Querier, categoryID string) (*domain.Category, error) {
	return r.get(ctx, q, `SELECT id, slug, name, description, active, sort_order, created_at, updated_at
		FROM categories WHERE id = ?`, categoryID)
}

func (r *Repo) GetBySlug(ctx context.Context, q db.Querier, slug string) (*domain.Category, error) {
	return r.get(ctx, q, `SELECT id, slug, name, description, active, sort_order, created_at, updated_at
		FROM categories WHERE slug = ?`, slug)
}

func (r *Repo) Search(ctx context.Context, q db.Querier, normalizedTerm string) ([]domain.Category, error) {
	like := "%" + normalizedTerm + "%"
	rows, err := q.QueryContext(ctx, `SELECT DISTINCT c.id, c.slug, c.name, c.description, c.active, c.sort_order, c.created_at, c.updated_at
		FROM categories c
		LEFT JOIN category_synonyms s ON s.category_id = c.id
		WHERE c.active = 1 AND (c.slug LIKE ? OR lower(c.name) LIKE ? OR s.normalized_term LIKE ?)
		ORDER BY c.sort_order, c.name LIMIT 20`, like, like, like)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []domain.Category
	for rows.Next() {
		var c domain.Category
		var active int
		var created, updated string
		if err := rows.Scan(&c.ID, &c.Slug, &c.Name, &c.Description, &active, &c.SortOrder, &created, &updated); err != nil {
			return nil, err
		}
		c.Active = active == 1
		c.CreatedAt, _ = time.Parse(time.RFC3339Nano, created)
		c.UpdatedAt, _ = time.Parse(time.RFC3339Nano, updated)
		c.Synonyms, _ = r.synonyms(ctx, q, c.ID)
		out = append(out, c)
	}
	return out, rows.Err()
}

func (r *Repo) list(ctx context.Context, q db.Querier, query string) ([]domain.Category, error) {
	rows, err := q.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []domain.Category
	for rows.Next() {
		var c domain.Category
		var active int
		var created, updated string
		if err := rows.Scan(&c.ID, &c.Slug, &c.Name, &c.Description, &active, &c.SortOrder, &created, &updated); err != nil {
			return nil, err
		}
		c.Active = active == 1
		c.CreatedAt, _ = time.Parse(time.RFC3339Nano, created)
		c.UpdatedAt, _ = time.Parse(time.RFC3339Nano, updated)
		c.Synonyms, _ = r.synonyms(ctx, q, c.ID)
		out = append(out, c)
	}
	return out, rows.Err()
}

func (r *Repo) synonyms(ctx context.Context, q db.Querier, categoryID string) ([]string, error) {
	rows, err := q.QueryContext(ctx, `SELECT term FROM category_synonyms WHERE category_id = ? ORDER BY term`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var term string
		if err := rows.Scan(&term); err != nil {
			return nil, err
		}
		out = append(out, term)
	}
	return out, rows.Err()
}

func (r *Repo) Create(ctx context.Context, q db.Querier, c domain.Category) error {
	_, err := q.ExecContext(ctx, `INSERT INTO categories (id, slug, name, description, active, sort_order, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, c.ID, c.Slug, c.Name, c.Description, boolInt(c.Active),
		c.SortOrder, formatTime(c.CreatedAt), formatTime(c.UpdatedAt))
	if err != nil && strings.Contains(err.Error(), "UNIQUE") {
		return app.ErrSlugTaken
	}
	return err
}

func (r *Repo) Update(ctx context.Context, q db.Querier, c domain.Category) error {
	_, err := q.ExecContext(ctx, `UPDATE categories SET name = ?, description = ?, sort_order = ?, updated_at = ?
		WHERE id = ?`, c.Name, c.Description, c.SortOrder, formatTime(c.UpdatedAt), c.ID)
	return err
}

func (r *Repo) ReplaceSynonyms(ctx context.Context, q db.Querier, categoryID string, synonyms []string, at time.Time) error {
	if _, err := q.ExecContext(ctx, `DELETE FROM category_synonyms WHERE category_id = ?`, categoryID); err != nil {
		return err
	}
	for _, term := range synonyms {
		normalized := domain.NormalizeTerm(term)
		if normalized == "" {
			continue
		}
		if _, err := q.ExecContext(ctx, `INSERT INTO category_synonyms (id, category_id, term, normalized_term, created_at)
			VALUES (?, ?, ?, ?, ?)`, id.New("syn"), categoryID, term, normalized, formatTime(at)); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) Deactivate(ctx context.Context, q db.Querier, categoryID string, at time.Time) error {
	_, err := q.ExecContext(ctx, `UPDATE categories SET active = 0, updated_at = ? WHERE id = ?`, formatTime(at), categoryID)
	return err
}

func boolInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

func formatTime(t time.Time) string { return t.UTC().Format(time.RFC3339Nano) }

var _ app.Repo = (*Repo)(nil)
