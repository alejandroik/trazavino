-- name: GetReception :one
SELECT *
FROM reception
WHERE id = ?
LIMIT 1;

-- name: AddReception :execresult
INSERT INTO reception (id, weight, sugar)
VALUES (?, ?, ?);