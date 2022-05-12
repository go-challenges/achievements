ALTER TABLE challenges ADD CONSTRAINT pk_id_challenges PRIMARY KEY (id);

CREATE TABLE users (
    id serial NOT NULL PRIMARY KEY,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL
);

INSERT INTO users (first_name, last_name) Values ('first user', 'first user');

CREATE TABLE users_challenges (
    id serial NOT NULL PRIMARY KEY,
    user_id bigint NOT NULL,
    challenge_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES users(id) DEFERRABLE INITIALLY DEFERRED,
    FOREIGN KEY (challenge_id)
        REFERENCES challenges(id) DEFERRABLE INITIALLY DEFERRED
);




