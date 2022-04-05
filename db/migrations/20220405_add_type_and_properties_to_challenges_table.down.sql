ALTER TABLE challenges DROP COLUMN type;

ALTER TABLE challenges DROP COLUMN properties;
DROP TYPE challenge_types IF EXISTS;