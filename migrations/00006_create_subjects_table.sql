-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS subjects (
    id SERIAL PRIMARY KEY,
    class VARCHAR(255) UNIQUE NOT NULL,
    subject1 VARCHAR(255) NOT NULL,
    subject2 VARCHAR(255) NOT NULL,
    subject3 VARCHAR(255) NOT NULL,
    subject4 VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
    
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subjects;

-- +goose StatementEnd