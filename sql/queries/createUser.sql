-- name: CreateUser :one
INSERT INTO users (UserID, FirstName, Email, HashedPassword, CreatedAt, LastEdit)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
RETURNING *;