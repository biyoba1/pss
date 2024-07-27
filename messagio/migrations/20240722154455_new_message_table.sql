-- +goose Up
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    text VARCHAR(255),
    processed BOOLEAN
);

-- +goose Down
DROP TABLE messages;
