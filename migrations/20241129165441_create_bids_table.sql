-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE bids ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE bids ADD COLUMN partner_id INTEGER NOT NULL;
ALTER TABLE bids ADD COLUMN account_id INTEGER NOT NULL;
ALTER TABLE bids ADD COLUMN created_at TIMESTAMP NOT NULL;
ALTER TABLE bids ADD COLUMN updated_at TIMESTAMP NOT NULL;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
