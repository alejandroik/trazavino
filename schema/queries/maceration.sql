-- name: GetMaceration :one
SELECT *
FROM maceration
WHERE id = $1
LIMIT 1;

-- name: ListMacerations :many
SELECT *
FROM maceration
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddMaceration :exec
INSERT INTO maceration (id, created_at, reception_id, warehouse_id)
VALUES ($1, $2, $3, $4);