ALTER TABLE users_challenges ADD COLUMN IF NOT EXISTS started_at timestamp;
ALTER TABLE users_challenges ADD COLUMN IF NOT EXISTS finished_at timestamp;
ALTER TABLE users_challenges ADD COLUMN IF NOT EXISTS ended_at timestamp;