-- name: CreateUser :one
INSERT INTO users (UserID, FirstName, Email, HashedPassword)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;