-- name: GetGrapeType :one
SELECT *
FROM grape_type
WHERE id = $1
LIMIT 1;

-- name: ListGrapeTypes :many
SELECT *
FROM grape_type
ORDER BY id
OFFSET $1 LIMIT $2;

-- name: AddGrapeType :one
INSERT INTO grape_type (created_at, name)
VALUES ($1, $2)
RETURNING *;