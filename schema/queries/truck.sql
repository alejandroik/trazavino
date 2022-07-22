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
INSERT INTO truck (id, created_at, name)
VALUES ($1, $2, $3);