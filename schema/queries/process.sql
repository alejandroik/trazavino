-- name: GetProcess :one
SELECT *
FROM process
WHERE id = ?
LIMIT 1;

-- name: AddProcess :execresult
INSERT INTO process (start_date, p_type)
VALUES (?, ?);