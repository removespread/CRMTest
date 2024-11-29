-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE contacts ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE contacts ADD COLUMN name VARCHAR(255) NOT NULL;
ALTER TABLE contacts ADD COLUMN email VARCHAR(255) UNIQUE NOT NULL;
ALTER TABLE contacts ADD COLUMN phone VARCHAR(255) NOT NULL;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
