-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS subjects (
    id SERIAL PRIMARY KEY,
    class INT NOT NULL,
    subject1 VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
    
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subjects;

-- +goose StatementEnd