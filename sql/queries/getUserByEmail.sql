-- name: GetUserByEmail :one
SELECT UserID, FirstName, Email, HashedPassword
FROM users 
WHERE Email = $1;