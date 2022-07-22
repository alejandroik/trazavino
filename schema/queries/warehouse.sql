-- name: GetWarehouse :one
SELECT *
FROM warehouse
WHERE id = $1
LIMIT 1;

-- name: ListWarehouses :many
SELECT *
FROM warehouse
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddWarehouse :exec
INSERT INTO warehouse (id, created_at, name, is_empty)
VALUES ($1, $2, $3, $4);

-- name: UpdateWarehouse :exec
UPDATE warehouse
SET updated_at = $2,
    is_empty   = $3
WHERE id = $1;