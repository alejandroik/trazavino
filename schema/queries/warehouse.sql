-- name: GetWarehouse :one
SELECT *
FROM warehouse
WHERE id = $1
LIMIT 1;

-- name: ListWarehouses :many
SELECT *
FROM warehouse
ORDER BY id
OFFSET $1 LIMIT $2;

-- name: AddWarehouse :one
INSERT INTO warehouse (created_at, name, is_empty)
VALUES ($1, $2, $3)
RETURNING *;