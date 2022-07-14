-- name: GetTruck :one
SELECT *
FROM truck
WHERE id = $1
LIMIT 1;

-- name: AddTruck :one
INSERT INTO truck (created_at, name)
VALUES ($1, $2)
RETURNING id;