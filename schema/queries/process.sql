-- name: GetProcess :one
SELECT *
FROM process
WHERE id = $1
LIMIT 1;

-- name: ListProcesses :many
SELECT *
FROM process
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddProcess :exec
INSERT INTO process (id, created_at, start_time, p_type)
VALUES ($1, $2, $3, $4);

-- name: UpdateProcess :exec
UPDATE process
SET updated_at  = $2,
    end_time    = $3,
    previous_id = $4,
    hash        = $5,
    transaction = $6
WHERE id = $1;