-- name: GetFermentation :one
SELECT *
FROM fermentation
WHERE id = $1
LIMIT 1;

-- name: FindFermentation :one
SELECT fermentation.*
FROM fermentation
         INNER JOIN process p on fermentation.id = p.id
WHERE end_time IS NULL
  AND fermentation.tank_id = $1
LIMIT 1;

-- name: ListFermentations :many
SELECT *
FROM fermentation
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddFermentation :exec
INSERT INTO fermentation (id, created_at, warehouse_id, tank_id)
VALUES ($1, $2, $3, $4);