CREATE TYPE challenge_types AS ENUM ('number', 'time');

ALTER TABLE challenges ADD COLUMN type challenge_types NOT NULL DEFAULT 'number'::challenge_types;

ALTER TABLE challenges ADD COLUMN properties jsonb NOT NULL DEFAULT '{}';