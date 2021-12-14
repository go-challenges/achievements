CREATE TABLE challanges (
    id serial NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    conditions jsonb
);