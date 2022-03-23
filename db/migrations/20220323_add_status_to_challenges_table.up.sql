CREATE TYPE challenge_statuses AS ENUM ('created', 'active', 'finished');

ALTER TABLE challenges ADD COLUMN status challenge_statuses NOT NULL DEFAULT 'created'::challenge_statuses;