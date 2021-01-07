-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE production_teams
(
    id         SERIAL PRIMARY KEY,
    name       TEXT,
    team_type  TEXT,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    hub_id     INTEGER REFERENCES hubs(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS production_teams CASCADE;