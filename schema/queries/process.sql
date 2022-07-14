-- name: GetProcess :one
SELECT *
FROM process
WHERE id = $1
LIMIT 1;

-- name: AddProcess :one
INSERT INTO process (created_at, start_date, p_type)
VALUES ($1, $2, $3)
RETURNING id;