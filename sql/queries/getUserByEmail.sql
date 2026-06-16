-- name: GetUserByEmail :one
SELECT UserID, FirstName, Email 
FROM users 
WHERE Email = $1;