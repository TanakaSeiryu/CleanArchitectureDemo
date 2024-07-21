-- +goose Up
-- +goose StatementBegin
CREATE TABLE grades (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		subject_id INT NOT NULL,
		grade INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE grades
-- +goose StatementEnd
