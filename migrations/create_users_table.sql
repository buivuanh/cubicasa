-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    email       TEXT,
    role       TEXT,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    team_id     INTEGER REFERENCES production_teams(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS users CASCADE;