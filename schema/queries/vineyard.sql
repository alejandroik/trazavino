-- name: GetVineyard :one
SELECT *
FROM vineyard
WHERE id = $1
LIMIT 1;

-- name: AddVineyard :one
INSERT INTO vineyard (created_at, name)
VALUES ($1, $2)
RETURNING id;