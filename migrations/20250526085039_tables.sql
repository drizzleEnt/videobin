-- +goose Up
CREATE TABLE files (
    id UUID PRIMARY KEY,
    file_name TEXT NOT NULL,
    file_type TEXT NOT NULL,
    size BIGINT NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE file_chunks (
    id SERIAL PRIMARY KEY,
    file_id UUID REFERENCES files(id),
    chunk_index INT NOT NULL,
    chunk_hash TEXT NOT NULL,
    object_name TEXT NOT NULL,
    size BIGINT NOT NULL
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
