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

-- name: UpdateTank :exec
UPDATE tank
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at),
    is_empty   = COALESCE($4, is_empty)
WHERE id = $1;