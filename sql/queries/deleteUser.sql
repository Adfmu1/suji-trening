-- name: DeleteUser :exec
DELETE FROM users
WHERE 
Email = $1 AND HashedPassword = $2;
