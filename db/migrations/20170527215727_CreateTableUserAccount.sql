
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS user_account (
	id bigserial PRIMARY KEY,
	user_name text NOT NULL,
	phone text NOT NULL,
	token text NOT NULL,
	created_on timestamp without time zone,
    modified_on timestamp without time zone
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS user_account;
