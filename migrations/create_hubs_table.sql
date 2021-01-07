-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE hubs
(
    id           SERIAL PRIMARY KEY,
    name         TEXT,
    created_at   timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at   timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    geo_location JSONB
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS hubs CASCADE;