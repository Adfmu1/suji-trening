-- +goose Up
CREATE TABLE users (
UserID TEXT PRIMARY KEY,
FirstName TEXT NOT NULL,
Email TEXT NOT NULL UNIQUE,
HashedPassword TEXT NOT NULL DEFAULT 'unset'
);

-- +goose Down
DROP TABLE users;