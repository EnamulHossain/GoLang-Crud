-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS marks (
		id SERIAL PRIMARY KEY,
		student_id INT REFERENCES students(id) ON DELETE CASCADE,
		datastructures INT NOT NULL,
		algorithms INT NOT NULL,
		computernetworks INT NOT NULL,
		artificialintelligence INT NOT NULL,
		operatingsystems INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP DEFAULT NULL
	  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE marks IF EXISTS;
-- +goose StatementEnd
