package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/raire/raire/internal/identity/app"
	"github.com/raire/raire/internal/identity/domain"
	"github.com/raire/raire/internal/platform/db"
	"github.com/raire/raire/internal/platform/id"
)

type Repo struct {
	DB *sql.DB
}

func New(database *sql.DB) *Repo { return &Repo{DB: database} }

const accountColumns = `id, email_normalized, password_hash, display_name, role, email_verified_at,
	contact_email, contact_phone, contact_phone_verified_at, contact_preference,
	mfa_secret, mfa_recovery_codes_hash, suspended_at, deleted_at, created_at, updated_at`

func (r *Repo) CreateAccount(ctx context.Context, q db.Querier, a domain.Account) error {
	recovery, _ := json.Marshal(a.MFARecoveryCodesHash)
	_, err := q.ExecContext(ctx, `INSERT INTO accounts (`+accountColumns+`)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		a.ID, a.EmailNormalized, a.PasswordHash, a.DisplayName, string(a.Role),
		timePtr(a.EmailVerifiedAt), a.ContactEmail, a.ContactPhone, timePtr(a.ContactPhoneVerifiedAt),
		string(a.ContactPreference), nullString(a.MFASecret), string(recovery),
		timePtr(a.SuspendedAt), timePtr(a.DeletedAt), formatTime(a.CreatedAt), formatTime(a.UpdatedAt))
	return err
}

func (r *Repo) GetAccountByEmail(ctx context.Context, q db.Querier, email string) (*domain.Account, error) {
	return r.getAccount(ctx, q, `SELECT `+accountColumns+` FROM accounts WHERE email_normalized = ?`, email)
}

func (r *Repo) GetAccountByID(ctx context.Context, q db.Querier, id string) (*domain.Account, error) {
	return r.getAccount(ctx, q, `SELECT `+accountColumns+` FROM accounts WHERE id = ?`, id)
}

func (r *Repo) getAccount(ctx context.Context, q db.Querier, query string, arg any) (*domain.Account, error) {
	var a domain.Account
	var role, pref, recovery string
	var emailVerified, phoneVerified, suspended, deleted sql.NullString
	var mfa sql.NullString
	var created, updated string
	err := q.QueryRowContext(ctx, query, arg).Scan(
		&a.ID, &a.EmailNormalized, &a.PasswordHash, &a.DisplayName, &role, &emailVerified,
		&a.ContactEmail, &a.ContactPhone, &phoneVerified, &pref,
		&mfa, &recovery, &suspended, &deleted, &created, &updated,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, app.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	a.Role = domain.Role(role)
	a.ContactPreference = domain.ContactPreference(pref)
	a.EmailVerifiedAt = parseTimePtr(emailVerified)
	a.ContactPhoneVerifiedAt = parseTimePtr(phoneVerified)
	a.SuspendedAt = parseTimePtr(suspended)
	a.DeletedAt = parseTimePtr(deleted)
	a.MFASecret = mfa.String
	_ = json.Unmarshal([]byte(recovery), &a.MFARecoveryCodesHash)
	a.CreatedAt = parseTime(created)
	a.UpdatedAt = parseTime(updated)
	return &a, nil
}

func (r *Repo) UpdateAccount(ctx context.Context, q db.Querier, a domain.Account) error {
	recovery, _ := json.Marshal(a.MFARecoveryCodesHash)
	_, err := q.ExecContext(ctx, `UPDATE accounts SET
		email_normalized = ?, password_hash = ?, display_name = ?, role = ?, email_verified_at = ?,
		contact_email = ?, contact_phone = ?, contact_phone_verified_at = ?, contact_preference = ?,
		mfa_secret = ?, mfa_recovery_codes_hash = ?, suspended_at = ?, deleted_at = ?, updated_at = ?
		WHERE id = ?`,
		a.EmailNormalized, a.PasswordHash, a.DisplayName, string(a.Role), timePtr(a.EmailVerifiedAt),
		a.ContactEmail, a.ContactPhone, timePtr(a.ContactPhoneVerifiedAt), string(a.ContactPreference),
		nullString(a.MFASecret), string(recovery), timePtr(a.SuspendedAt), timePtr(a.DeletedAt),
		formatTime(a.UpdatedAt), a.ID)
	return err
}

func (r *Repo) SetAccountRole(ctx context.Context, q db.Querier, id string, role domain.Role) error {
	_, err := q.ExecContext(ctx, `UPDATE accounts SET role = ?, updated_at = ? WHERE id = ?`,
		string(role), formatTime(time.Now().UTC()), id)
	return err
}

func (r *Repo) CreateConsent(ctx context.Context, q db.Querier, c domain.Consent) error {
	_, err := q.ExecContext(ctx, `INSERT INTO consents (id, account_id, kind, version, granted_at, revoked_at)
		VALUES (?, ?, ?, ?, ?, ?)`, c.ID, c.AccountID, string(c.Kind), c.Version,
		formatTime(c.GrantedAt), timePtr(c.RevokedAt))
	return err
}

func (r *Repo) ListConsents(ctx context.Context, q db.Querier, accountID string) ([]domain.Consent, error) {
	rows, err := q.QueryContext(ctx, `SELECT id, account_id, kind, version, granted_at, revoked_at
		FROM consents WHERE account_id = ? ORDER BY granted_at`, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []domain.Consent
	for rows.Next() {
		var c domain.Consent
		var kind, granted string
		var revoked sql.NullString
		if err := rows.Scan(&c.ID, &c.AccountID, &kind, &c.Version, &granted, &revoked); err != nil {
			return nil, err
		}
		c.Kind = domain.ConsentKind(kind)
		c.GrantedAt = parseTime(granted)
		c.RevokedAt = parseTimePtr(revoked)
		out = append(out, c)
	}
	return out, rows.Err()
}

func (r *Repo) RevokeConsent(ctx context.Context, q db.Querier, accountID string, kind domain.ConsentKind, at time.Time) error {
	res, err := q.ExecContext(ctx, `UPDATE consents SET revoked_at = ?
		WHERE id = (SELECT id FROM consents WHERE account_id = ? AND kind = ? AND revoked_at IS NULL
		ORDER BY granted_at DESC LIMIT 1)`, formatTime(at), accountID, string(kind))
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return app.ErrNotFound
	}
	return nil
}

func (r *Repo) CreateEmailToken(ctx context.Context, q db.Querier, t domain.EmailToken) error {
	_, err := q.ExecContext(ctx, `INSERT INTO email_tokens (id, account_id, kind, token_hash, expires_at, used_at, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`, t.ID, t.AccountID, t.Kind, t.TokenHash,
		formatTime(t.ExpiresAt), timePtr(t.UsedAt), formatTime(t.CreatedAt))
	return err
}

func (r *Repo) GetEmailTokenByHash(ctx context.Context, q db.Querier, hash string) (*domain.EmailToken, error) {
	var t domain.EmailToken
	var expires, created string
	var used sql.NullString
	err := q.QueryRowContext(ctx, `SELECT id, account_id, kind, token_hash, expires_at, used_at, created_at
		FROM email_tokens WHERE token_hash = ?`, hash).Scan(&t.ID, &t.AccountID, &t.Kind, &t.TokenHash,
		&expires, &used, &created)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, app.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	t.ExpiresAt = parseTime(expires)
	t.CreatedAt = parseTime(created)
	t.UsedAt = parseTimePtr(used)
	return &t, nil
}

func (r *Repo) MarkEmailTokenUsed(ctx context.Context, q db.Querier, id string, at time.Time) (bool, error) {
	res, err := q.ExecContext(ctx, `UPDATE email_tokens SET used_at = ? WHERE id = ? AND used_at IS NULL`, formatTime(at), id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n == 1, nil
}

func (r *Repo) InvalidateEmailTokens(ctx context.Context, q db.Querier, accountID, kind string, at time.Time) error {
	_, err := q.ExecContext(ctx, `UPDATE email_tokens SET used_at = ?
		WHERE account_id = ? AND kind = ? AND used_at IS NULL`, formatTime(at), accountID, kind)
	return err
}

func (r *Repo) CreateSession(ctx context.Context, q db.Querier, s domain.Session) error {
	_, err := q.ExecContext(ctx, `INSERT INTO sessions
		(id, account_id, token_hash, csrf_token_hash, mfa_verified_at, created_at, expires_at, revoked_at, last_seen_at, ip_hash, user_agent_hash)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		s.ID, s.AccountID, s.TokenHash, s.CSRFTokenHash, timePtr(s.MFAVerifiedAt),
		formatTime(s.CreatedAt), formatTime(s.ExpiresAt), timePtr(s.RevokedAt),
		timePtr(s.LastSeenAt), s.IPHash, s.UserAgentHash)
	return err
}

func (r *Repo) GetSessionByTokenHash(ctx context.Context, q db.Querier, hash string) (*domain.Session, error) {
	var s domain.Session
	var mfa, revoked, lastSeen sql.NullString
	var created, expires string
	err := q.QueryRowContext(ctx, `SELECT id, account_id, token_hash, csrf_token_hash, mfa_verified_at,
		created_at, expires_at, revoked_at, last_seen_at, ip_hash, user_agent_hash
		FROM sessions WHERE token_hash = ?`, hash).Scan(&s.ID, &s.AccountID, &s.TokenHash, &s.CSRFTokenHash,
		&mfa, &created, &expires, &revoked, &lastSeen, &s.IPHash, &s.UserAgentHash)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, app.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	s.MFAVerifiedAt = parseTimePtr(mfa)
	s.CreatedAt = parseTime(created)
	s.ExpiresAt = parseTime(expires)
	s.RevokedAt = parseTimePtr(revoked)
	s.LastSeenAt = parseTimePtr(lastSeen)
	return &s, nil
}

func (r *Repo) RevokeSessionByHash(ctx context.Context, q db.Querier, hash string, at time.Time) error {
	_, err := q.ExecContext(ctx, `UPDATE sessions SET revoked_at = ? WHERE token_hash = ? AND revoked_at IS NULL`,
		formatTime(at), hash)
	return err
}

func (r *Repo) RevokeSessionByID(ctx context.Context, q db.Querier, id string, at time.Time) error {
	_, err := q.ExecContext(ctx, `UPDATE sessions SET revoked_at = ? WHERE id = ? AND revoked_at IS NULL`,
		formatTime(at), id)
	return err
}

func (r *Repo) RevokeOtherSessions(ctx context.Context, q db.Querier, accountID, keepSessionID string, at time.Time) error {
	_, err := q.ExecContext(ctx, `UPDATE sessions SET revoked_at = ?
		WHERE account_id = ? AND revoked_at IS NULL AND id <> ?`, formatTime(at), accountID, keepSessionID)
	return err
}

func (r *Repo) RevokeAllSessions(ctx context.Context, q db.Querier, accountID string, at time.Time) error {
	_, err := q.ExecContext(ctx, `UPDATE sessions SET revoked_at = ? WHERE account_id = ? AND revoked_at IS NULL`,
		formatTime(at), accountID)
	return err
}

func (r *Repo) TouchSession(ctx context.Context, q db.Querier, id string, at time.Time) error {
	_, err := q.ExecContext(ctx, `UPDATE sessions SET last_seen_at = ? WHERE id = ?`, formatTime(at), id)
	return err
}

func (r *Repo) UpsertVerification(ctx context.Context, q db.Querier, accountID, kind, status string, at time.Time) error {
	_, err := q.ExecContext(ctx, `INSERT INTO verifications (id, account_id, kind, status, decided_at, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(account_id, kind) DO UPDATE SET status = excluded.status, decided_at = excluded.decided_at`,
		"ver_"+accountID+"_"+kind, accountID, kind, status, formatTime(at), formatTime(at))
	return err
}

func (r *Repo) ReplaceRecoveryCodes(ctx context.Context, q db.Querier, accountID string, hashes []string, at time.Time) error {
	if _, err := q.ExecContext(ctx, `DELETE FROM mfa_recovery_codes WHERE account_id = ?`, accountID); err != nil {
		return err
	}
	for _, hash := range hashes {
		if _, err := q.ExecContext(ctx, `INSERT INTO mfa_recovery_codes (id, account_id, code_hash, created_at)
			VALUES (?, ?, ?, ?)`, id.New("mfa"), accountID, hash, formatTime(at)); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) ConsumeRecoveryCode(ctx context.Context, q db.Querier, accountID, hash string, at time.Time) (bool, error) {
	res, err := q.ExecContext(ctx, `UPDATE mfa_recovery_codes SET used_at = ?
		WHERE account_id = ? AND code_hash = ? AND used_at IS NULL`, formatTime(at), accountID, hash)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n == 1, nil
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

func nullString(v string) any {
	if v == "" {
		return nil
	}
	return v
}

var _ app.Repo = (*Repo)(nil)

func (r *Repo) String() string { return fmt.Sprintf("identity.sqlite(%p)", r.DB) }
