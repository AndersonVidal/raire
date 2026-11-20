package sqlite

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/raire/raire/internal/directory/app"
	"github.com/raire/raire/internal/directory/domain"
	"github.com/raire/raire/internal/platform/db"
)

type Repo struct {
	DB *sql.DB
}

func New(database *sql.DB) *Repo { return &Repo{DB: database} }

const profileColumns = `id, account_id, display_name, bio, service_type, status, city, neighborhood,
	public_area_label, precise_lat_enc, precise_lng_enc, precise_location_expires_at, radius_km,
	availability_windows_json, availability_note, contact_preference_override, verification,
	published_at, created_at, updated_at, suspended_at, suspended_reason, version`

func (r *Repo) Create(ctx context.Context, q db.Querier, p domain.Profile) error {
	availability, _ := json.Marshal(p.Availability)
	if p.Version == 0 {
		p.Version = 1
	}
	_, err := q.ExecContext(ctx, `INSERT INTO professional_profiles (`+profileColumns+`)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.ID, p.AccountID, p.DisplayName, p.Bio, string(p.ServiceType), string(p.Status),
		p.City, p.Neighborhood, p.PublicAreaLabel, nullString(p.PreciseLatEnc), nullString(p.PreciseLngEnc),
		timePtr(p.PreciseLocationExpiresAt), intPtr(p.RadiusKm), string(availability), p.AvailabilityNote,
		prefPtr(p.ContactPreferenceOverride), string(p.Verification), timePtr(p.PublishedAt),
		formatTime(p.CreatedAt), formatTime(p.UpdatedAt), timePtr(p.SuspendedAt), p.SuspendedReason, p.Version)
	if err != nil && strings.Contains(err.Error(), "UNIQUE") {
		return app.ErrProfileExists
	}
	return err
}

func (r *Repo) Update(ctx context.Context, q db.Querier, p domain.Profile) error {
	availability, _ := json.Marshal(p.Availability)
	res, err := q.ExecContext(ctx, `UPDATE professional_profiles SET
		display_name = ?, bio = ?, service_type = ?, status = ?, city = ?, neighborhood = ?, public_area_label = ?,
		precise_lat_enc = ?, precise_lng_enc = ?, precise_location_expires_at = ?, radius_km = ?,
		availability_windows_json = ?, availability_note = ?, contact_preference_override = ?, verification = ?,
		published_at = ?, updated_at = ?, suspended_at = ?, suspended_reason = ?, version = version + 1
		WHERE id = ? AND version = ?`,
		p.DisplayName, p.Bio, string(p.ServiceType), string(p.Status), p.City, p.Neighborhood, p.PublicAreaLabel,
		nullString(p.PreciseLatEnc), nullString(p.PreciseLngEnc), timePtr(p.PreciseLocationExpiresAt),
		intPtr(p.RadiusKm), string(availability), p.AvailabilityNote, prefPtr(p.ContactPreferenceOverride),
		string(p.Verification), timePtr(p.PublishedAt), formatTime(p.UpdatedAt), timePtr(p.SuspendedAt),
		p.SuspendedReason, p.ID, p.Version)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return app.ErrConflict
	}
	return nil
}

func (r *Repo) GetByAccountID(ctx context.Context, q db.Querier, accountID string) (*domain.Profile, error) {
	return r.get(ctx, q, `SELECT `+profileColumns+` FROM professional_profiles WHERE account_id = ?`, accountID)
}

func (r *Repo) GetByID(ctx context.Context, q db.Querier, profileID string) (*domain.Profile, error) {
	return r.get(ctx, q, `SELECT `+profileColumns+` FROM professional_profiles WHERE id = ?`, profileID)
}

func (r *Repo) get(ctx context.Context, q db.Querier, query string, arg any) (*domain.Profile, error) {
	var p domain.Profile
	var serviceType, status, verification, availability, created, updated string
	var latEnc, lngEnc, expires, override, published, suspended sql.NullString
	var radius sql.NullInt64
	err := q.QueryRowContext(ctx, query, arg).Scan(
		&p.ID, &p.AccountID, &p.DisplayName, &p.Bio, &serviceType, &status, &p.City, &p.Neighborhood,
		&p.PublicAreaLabel, &latEnc, &lngEnc, &expires, &radius,
		&availability, &p.AvailabilityNote, &override, &verification,
		&published, &created, &updated, &suspended, &p.SuspendedReason, &p.Version,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, app.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	p.ServiceType = domain.ServiceType(serviceType)
	p.Status = domain.Status(status)
	p.Verification = domain.Verification(verification)
	p.PreciseLatEnc, p.PreciseLngEnc = latEnc.String, lngEnc.String
	p.PreciseLocationExpiresAt = parseTimePtr(expires)
	if radius.Valid {
		v := int(radius.Int64)
		p.RadiusKm = &v
	}
	_ = json.Unmarshal([]byte(availability), &p.Availability)
	if override.Valid {
		pref := domain.ContactPreference(override.String)
		p.ContactPreferenceOverride = &pref
	}
	p.PublishedAt = parseTimePtr(published)
	p.CreatedAt = parseTime(created)
	p.UpdatedAt = parseTime(updated)
	p.SuspendedAt = parseTimePtr(suspended)
	p.CategoryIDs, err = r.categoryIDs(ctx, q, p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repo) categoryIDs(ctx context.Context, q db.Querier, profileID string) ([]string, error) {
	rows, err := q.QueryContext(ctx, `SELECT category_id FROM profile_categories WHERE profile_id = ? ORDER BY category_id`, profileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}

func (r *Repo) ReplaceCategories(ctx context.Context, q db.Querier, profileID string, categoryIDs []string, at time.Time) error {
	if _, err := q.ExecContext(ctx, `DELETE FROM profile_categories WHERE profile_id = ?`, profileID); err != nil {
		return err
	}
	for i, categoryID := range categoryIDs {
		primary := 0
		if i == 0 {
			primary = 1
		}
		if _, err := q.ExecContext(ctx, `INSERT INTO profile_categories (profile_id, category_id, primary_category)
			VALUES (?, ?, ?)`, profileID, categoryID, primary); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) List(ctx context.Context, q db.Querier, f app.ListFilter) (app.ListPage, error) {
	limit := f.Limit
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	where := []string{"1=1"}
	args := []any{}
	if f.Status != "" {
		where = append(where, "status = ?")
		args = append(args, f.Status)
	}
	if strings.TrimSpace(f.Query) != "" {
		where = append(where, "lower(display_name) LIKE ?")
		args = append(args, "%"+strings.ToLower(strings.TrimSpace(f.Query))+"%")
	}
	if f.Cursor != "" {
		raw, err := base64.RawURLEncoding.DecodeString(f.Cursor)
		if err != nil {
			return app.ListPage{}, errors.New("invalid cursor")
		}
		parts := strings.SplitN(string(raw), "|", 2)
		if len(parts) != 2 {
			return app.ListPage{}, errors.New("invalid cursor")
		}
		where = append(where, "(updated_at < ? OR (updated_at = ? AND id < ?))")
		args = append(args, parts[0], parts[0], parts[1])
	}
	args = append(args, limit+1)
	rows, err := q.QueryContext(ctx, `SELECT id, account_id, display_name, status, city, neighborhood, updated_at
		FROM professional_profiles WHERE `+strings.Join(where, " AND ")+`
		ORDER BY updated_at DESC, id DESC LIMIT ?`, args...)
	if err != nil {
		return app.ListPage{}, err
	}
	defer rows.Close()
	page := app.ListPage{}
	for rows.Next() {
		var s app.AdminSummary
		var status, updated string
		if err := rows.Scan(&s.ID, &s.AccountID, &s.DisplayName, &status, &s.City, &s.Neighborhood, &updated); err != nil {
			return app.ListPage{}, err
		}
		s.Status = status
		s.UpdatedAt = parseTime(updated)
		page.Items = append(page.Items, s)
	}
	if err := rows.Err(); err != nil {
		return app.ListPage{}, err
	}
	if len(page.Items) > limit {
		last := page.Items[limit-1]
		page.NextCursor = base64.RawURLEncoding.EncodeToString([]byte(last.UpdatedAt.UTC().Format(time.RFC3339Nano) + "|" + last.ID))
		page.Items = page.Items[:limit]
	}
	return page, nil
}

func (r *Repo) PurgeExpiredLocations(ctx context.Context, q db.Querier, now time.Time) ([]string, error) {
	rows, err := q.QueryContext(ctx, `SELECT id FROM professional_profiles
		WHERE precise_lat_enc IS NOT NULL AND (precise_location_expires_at IS NULL OR precise_location_expires_at <= ?)`,
		now.UTC().Format(time.RFC3339Nano))
	if err != nil {
		return nil, err
	}
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			rows.Close()
			return nil, err
		}
		ids = append(ids, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	for _, id := range ids {
		if _, err := q.ExecContext(ctx, `UPDATE professional_profiles
			SET precise_lat_enc = NULL, precise_lng_enc = NULL, precise_location_expires_at = NULL, updated_at = ?
			WHERE id = ?`, now.UTC().Format(time.RFC3339Nano), id); err != nil {
			return nil, err
		}
	}
	return ids, nil
}

func formatTime(t time.Time) string { return t.UTC().Format(time.RFC3339Nano) }

func parseTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339Nano, s)
	return t
}

func parseTimePtr(s sql.NullString) *time.Time {
	if !s.Valid || s.String == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, s.String)
	if err != nil {
		return nil
	}
	return &t
}

func timePtr(t *time.Time) any {
	if t == nil {
		return nil
	}
	return t.UTC().Format(time.RFC3339Nano)
}

func intPtr(v *int) any {
	if v == nil {
		return nil
	}
	return *v
}

func prefPtr(p *domain.ContactPreference) any {
	if p == nil {
		return nil
	}
	return string(*p)
}

func nullString(v string) any {
	if v == "" {
		return nil
	}
	return v
}

var _ app.Repo = (*Repo)(nil)
