-- 0003_profile_version.sql — optimistic concurrency for profile mutations.
-- Prevents an owner PATCH from overwriting a concurrent administrative
-- suspension (lost update).
ALTER TABLE professional_profiles ADD COLUMN version INTEGER NOT NULL DEFAULT 1;
