-- 0004_mfa_recovery_codes.sql — one-time recovery codes with atomic consumption.
CREATE TABLE IF NOT EXISTS mfa_recovery_codes (
    id         TEXT PRIMARY KEY,
    account_id TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    code_hash  TEXT NOT NULL,
    used_at    TEXT,
    created_at TEXT NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS ux_mfa_recovery_code ON mfa_recovery_codes(account_id, code_hash);
CREATE INDEX IF NOT EXISTS ix_mfa_recovery_account ON mfa_recovery_codes(account_id, used_at);
