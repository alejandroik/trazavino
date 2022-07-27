-- name: GetReception :one
SELECT *
FROM reception
WHERE id = $1
LIMIT 1;

-- name: ListReceptions :many
SELECT *
FROM reception
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddReception :exec
INSERT INTO reception (id, created_at, weight, sugar, truck_id, vineyard_id, grape_type_id)
VALUES ($1, $2, $3, $4, $5, $6, $7);