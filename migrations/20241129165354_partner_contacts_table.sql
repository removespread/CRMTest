-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE partners ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE partners ADD COLUMN name VARCHAR(255) NOT NULL;
ALTER TABLE partners ADD COLUMN contacts JSONB;
ALTER TABLE partners ADD COLUMN description TEXT NOT NULL;


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
