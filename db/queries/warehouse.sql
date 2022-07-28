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
INSERT INTO warehouse (id, created_at, winery_id, name, is_empty)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateWarehouseUsage :exec
UPDATE warehouse
SET is_empty   = $2,
    updated_at = $3
WHERE id = $1;

-- name: UpdateWarehouseData :exec
UPDATE warehouse
SET name       = $2,
    updated_at = $3
WHERE id = $1;