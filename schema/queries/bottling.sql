-- name: GetBottling :one
SELECT *
FROM bottling
WHERE id = $1
LIMIT 1;

-- name: ListBottlings :many
SELECT *
FROM bottling
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddBottling :exec
INSERT INTO bottling (id, created_at, cask_id, wine_id, bottle_qty)
VALUES ($1, $2, $3, $4, $5);