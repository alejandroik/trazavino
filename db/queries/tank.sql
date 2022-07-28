-- name: GetTank :one
SELECT *
FROM tank
WHERE id = $1
LIMIT 1;

-- name: ListTanks :many
SELECT *
FROM tank
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddTank :exec
INSERT INTO tank (id, created_at, winery_id, name, is_empty)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateTankUsage :exec
UPDATE tank
SET is_empty   = $2,
    updated_at = $3
WHERE id = $1;

-- name: UpdateTankData :exec
UPDATE tank
SET name       = $2,
    updated_at = $3
WHERE id = $1;