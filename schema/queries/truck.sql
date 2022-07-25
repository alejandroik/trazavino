-- name: GetTruck :one
SELECT *
FROM truck
WHERE id = $1
LIMIT 1;

-- name: ListTrucks :many
SELECT *
FROM truck
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddTruck :exec
INSERT INTO truck (id, created_at, winery_id, name)
VALUES ($1, $2, $3, $4);

-- name: UpdateTruck :exec
UPDATE truck
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at)
WHERE id = $1;