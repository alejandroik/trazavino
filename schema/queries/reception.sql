-- name: GetReception :one
SELECT *
FROM reception
WHERE id = $1
LIMIT 1;

-- name: AddReception :execresult
INSERT INTO reception (id, created_at, weight, sugar)
VALUES ($1, $2, $3, $4);