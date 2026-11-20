-- 0001_init.sql — foundation-platform schema.
-- Forward-only. Rollback is logical (feature flag / previous binary); no
-- destructive down migration is shipped.

CREATE TABLE IF NOT EXISTS accounts (
    id                       TEXT PRIMARY KEY,
    email_normalized         TEXT NOT NULL,
    password_hash            TEXT NOT NULL,
    display_name             TEXT NOT NULL,
    role                     TEXT NOT NULL CHECK (role IN ('client','professional','operator','admin')),
    email_verified_at        TEXT,
    contact_email            TEXT,
    contact_phone            TEXT,
    contact_phone_verified_at TEXT,
    contact_preference       TEXT NOT NULL DEFAULT 'platform' CHECK (contact_preference IN ('platform','email','phone')),
    mfa_secret               TEXT,
    mfa_recovery_codes_hash  TEXT NOT NULL DEFAULT '[]',
    suspended_at             TEXT,
    deleted_at               TEXT,
    created_at               TEXT NOT NULL,
    updated_at               TEXT NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_accounts_email ON accounts(email_normalized);
CREATE INDEX IF NOT EXISTS ix_accounts_role ON accounts(role);

CREATE TABLE IF NOT EXISTS sessions (
    id              TEXT PRIMARY KEY,
    account_id      TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    token_hash      TEXT NOT NULL,
    csrf_token_hash TEXT NOT NULL,
    mfa_verified_at TEXT,
    created_at      TEXT NOT NULL,
    expires_at      TEXT NOT NULL,
    revoked_at      TEXT,
    last_seen_at    TEXT,
    ip_hash         TEXT,
    user_agent_hash TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_sessions_token ON sessions(token_hash);
CREATE INDEX IF NOT EXISTS ix_sessions_account ON sessions(account_id);

CREATE TABLE IF NOT EXISTS email_tokens (
    id         TEXT PRIMARY KEY,
    account_id TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    kind       TEXT NOT NULL CHECK (kind IN ('confirm','reset')),
    token_hash TEXT NOT NULL,
    expires_at TEXT NOT NULL,
    used_at    TEXT,
    created_at TEXT NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_email_tokens_hash ON email_tokens(token_hash);
CREATE INDEX IF NOT EXISTS ix_email_tokens_account ON email_tokens(account_id, kind);

CREATE TABLE IF NOT EXISTS consents (
    id         TEXT PRIMARY KEY,
    account_id TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    kind       TEXT NOT NULL CHECK (kind IN ('terms','privacy','research','marketing')),
    version    TEXT NOT NULL,
    granted_at TEXT NOT NULL,
    revoked_at TEXT
);
CREATE INDEX IF NOT EXISTS ix_consents_account ON consents(account_id, kind);

CREATE TABLE IF NOT EXISTS account_requests (
    id              TEXT PRIMARY KEY,
    account_id      TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    kind            TEXT NOT NULL CHECK (kind IN ('export','delete','correct')),
    status          TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending','processing','resolved','rejected')),
    requested_at    TEXT NOT NULL,
    resolved_at     TEXT,
    resolution_note TEXT
);
CREATE INDEX IF NOT EXISTS ix_account_requests_status ON account_requests(status, requested_at);

CREATE TABLE IF NOT EXISTS categories (
    id          TEXT PRIMARY KEY,
    slug        TEXT NOT NULL,
    name        TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    active      INTEGER NOT NULL DEFAULT 1 CHECK (active IN (0,1)),
    sort_order  INTEGER NOT NULL DEFAULT 0,
    created_at  TEXT NOT NULL,
    updated_at  TEXT NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_categories_slug ON categories(slug);
CREATE INDEX IF NOT EXISTS ix_categories_active ON categories(active, sort_order);

CREATE TABLE IF NOT EXISTS category_synonyms (
    id              TEXT PRIMARY KEY,
    category_id     TEXT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    term            TEXT NOT NULL,
    normalized_term TEXT NOT NULL,
    created_at      TEXT NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_category_synonyms ON category_synonyms(category_id, normalized_term);
CREATE INDEX IF NOT EXISTS ix_category_synonyms_normalized ON category_synonyms(normalized_term);

CREATE TABLE IF NOT EXISTS professional_profiles (
    id                            TEXT PRIMARY KEY,
    account_id                    TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    display_name                  TEXT NOT NULL,
    bio                           TEXT NOT NULL DEFAULT '',
    service_type                  TEXT NOT NULL CHECK (service_type IN ('in_person','remote','hybrid')),
    status                        TEXT NOT NULL DEFAULT 'draft' CHECK (status IN ('draft','published','paused','suspended')),
    city                          TEXT NOT NULL DEFAULT '',
    neighborhood                  TEXT NOT NULL DEFAULT '',
    public_area_label             TEXT NOT NULL DEFAULT '',
    precise_lat_enc               TEXT,
    precise_lng_enc               TEXT,
    precise_location_expires_at   TEXT,
    radius_km                     INTEGER,
    availability_windows_json     TEXT NOT NULL DEFAULT '[]',
    availability_note             TEXT NOT NULL DEFAULT '',
    contact_preference_override   TEXT CHECK (contact_preference_override IS NULL OR contact_preference_override IN ('platform','email','phone')),
    verification                  TEXT NOT NULL DEFAULT 'unverified' CHECK (verification IN ('unverified','contact_verified','identity_verified')),
    published_at                  TEXT,
    created_at                    TEXT NOT NULL,
    updated_at                    TEXT NOT NULL,
    suspended_at                  TEXT,
    suspended_reason              TEXT NOT NULL DEFAULT ''
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_profiles_account ON professional_profiles(account_id);
CREATE INDEX IF NOT EXISTS ix_profiles_status ON professional_profiles(status, updated_at);

CREATE TABLE IF NOT EXISTS profile_categories (
    profile_id  TEXT NOT NULL REFERENCES professional_profiles(id) ON DELETE CASCADE,
    category_id TEXT NOT NULL REFERENCES categories(id),
    primary_category INTEGER NOT NULL DEFAULT 0 CHECK (primary_category IN (0,1)),
    PRIMARY KEY (profile_id, category_id)
);
CREATE INDEX IF NOT EXISTS ix_profile_categories_category ON profile_categories(category_id);

CREATE TABLE IF NOT EXISTS portfolio_items (
    id             TEXT PRIMARY KEY,
    profile_id     TEXT NOT NULL REFERENCES professional_profiles(id) ON DELETE CASCADE,
    title          TEXT NOT NULL,
    description    TEXT NOT NULL DEFAULT '',
    category_id    TEXT NOT NULL REFERENCES categories(id),
    image_ref      TEXT,
    image_mime     TEXT,
    image_size_bytes INTEGER,
    position       INTEGER NOT NULL DEFAULT 0,
    status         TEXT NOT NULL DEFAULT 'draft' CHECK (status IN ('draft','published','removed')),
    created_at     TEXT NOT NULL,
    updated_at     TEXT NOT NULL,
    removed_at     TEXT
);
CREATE INDEX IF NOT EXISTS ix_portfolio_profile ON portfolio_items(profile_id, status, position);

CREATE TABLE IF NOT EXISTS verifications (
    id               TEXT PRIMARY KEY,
    account_id       TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    kind             TEXT NOT NULL CHECK (kind IN ('contact_email','contact_phone','identity')),
    status           TEXT NOT NULL CHECK (status IN ('not_started','pending','verified','rejected','expired')),
    evidence_ref_enc TEXT,
    decided_by       TEXT,
    decided_at       TEXT,
    expires_at       TEXT,
    created_at       TEXT NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_verifications_account_kind ON verifications(account_id, kind);

CREATE TABLE IF NOT EXISTS audit_events (
    id             TEXT PRIMARY KEY,
    actor_id       TEXT,
    actor_role     TEXT NOT NULL DEFAULT '',
    action         TEXT NOT NULL,
    target_type    TEXT NOT NULL,
    target_id      TEXT NOT NULL,
    reason         TEXT NOT NULL DEFAULT '',
    correlation_id TEXT NOT NULL DEFAULT '',
    metadata_json  TEXT NOT NULL DEFAULT '{}',
    created_at     TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS ix_audit_target ON audit_events(target_type, target_id, created_at);
CREATE INDEX IF NOT EXISTS ix_audit_actor ON audit_events(actor_id, created_at);

CREATE TRIGGER IF NOT EXISTS trg_audit_events_no_update
BEFORE UPDATE ON audit_events
BEGIN
    SELECT RAISE(ABORT, 'audit_events is append-only');
END;

CREATE TRIGGER IF NOT EXISTS trg_audit_events_no_delete
BEFORE DELETE ON audit_events
BEGIN
    SELECT RAISE(ABORT, 'audit_events is append-only');
END;

CREATE TABLE IF NOT EXISTS outbox_events (
    id           TEXT PRIMARY KEY,
    event_type   TEXT NOT NULL,
    payload_json TEXT NOT NULL DEFAULT '{}',
    occurred_at  TEXT NOT NULL,
    processed_at TEXT,
    attempts     INTEGER NOT NULL DEFAULT 0,
    last_error   TEXT
);
CREATE INDEX IF NOT EXISTS ix_outbox_pending ON outbox_events(processed_at, occurred_at);

CREATE TABLE IF NOT EXISTS feature_flags (
    key         TEXT PRIMARY KEY,
    enabled     INTEGER NOT NULL DEFAULT 0 CHECK (enabled IN (0,1)),
    description TEXT NOT NULL DEFAULT '',
    owner       TEXT NOT NULL DEFAULT '',
    review_at   TEXT,
    updated_at  TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS idempotency_keys (
    key          TEXT NOT NULL,
    operation    TEXT NOT NULL,
    request_hash TEXT NOT NULL,
    account_id   TEXT,
    created_at   TEXT NOT NULL,
    expires_at   TEXT NOT NULL,
    PRIMARY KEY (key, operation)
);
CREATE INDEX IF NOT EXISTS ix_idempotency_expires ON idempotency_keys(expires_at);
