-- name: GetUserByID :one
SELECT id, name, email, created_at
FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET
    name = $2
WHERE id = $1
RETURNING id, name, email, created_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;