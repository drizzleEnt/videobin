-- +goose Up
CREATE TABLE files (
    lint TEXT not null,
    chunk INTEGER NOT NULL DEFAULT 0
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd