-- name: CreateCategory :one
INSERT INTO categories (
    user_id,
    name,
    type,
    color,
    icon
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, name, type, color, icon, created_at;

-- name: GetCategoriesByUserID :many
SELECT id, user_id, name, type, color, icon, created_at
FROM categories
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetCategories :many
SELECT *
FROM categories
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateCategory :one
UPDATE categories
SET
    name = $2,
    type = $3,
    color = $4,
    icon = $5
WHERE id = $1 AND user_id = $6
RETURNING id, user_id, name, type, color, icon, created_at;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1 AND user_id = $2; 