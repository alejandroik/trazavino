-- name: GetGrapeType :one
SELECT *
FROM grape_type
WHERE id = $1
LIMIT 1;

-- name: ListGrapeTypes :many
SELECT *
FROM grape_type
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddGrapeType :exec
INSERT INTO grape_type (id, created_at, name)
VALUES ($1, $2, $3);