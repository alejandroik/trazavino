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

-- name: UpdateWarehouse :exec
UPDATE warehouse
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at),
    is_empty   = COALESCE($4, is_empty)
WHERE id = $1;