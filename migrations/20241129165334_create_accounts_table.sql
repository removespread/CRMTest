-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE accounts ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE accounts ADD COLUMN email VARCHAR(255) UNIQUE NOT NULL;
ALTER TABLE accounts ADD COLUMN password VARCHAR(255) NOT NULL;


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
