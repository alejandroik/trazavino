-- name: GetMaceration :one
SELECT *
FROM maceration
WHERE id = $1
LIMIT 1;

-- name: AddMaceration :one
INSERT INTO maceration (id, created_at, reception_id, warehouse_id)
VALUES ($1, $2, $3, $4)
RETURNING *;