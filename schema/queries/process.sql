-- name: GetProcess :one
SELECT *
FROM process
WHERE id = $1
LIMIT 1;

-- name: ListProcesses :many
SELECT *
FROM process
ORDER BY id DESC
OFFSET $1 LIMIT $2;

-- name: AddProcess :one
INSERT INTO process (created_at, start_date, p_type)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProcess :exec
UPDATE process
SET updated_at  = $2,
    end_date    = $3,
    previous_id = $4,
    hash        = $5,
    transaction = $6
WHERE id = $1;