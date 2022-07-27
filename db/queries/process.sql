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
INSERT INTO process (id, created_at, winery_id, start_time, p_type, previous_id)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateProcess :exec
UPDATE process
SET updated_at  = COALESCE($2, updated_at),
    end_time    = COALESCE($3, end_time),
    previous_id = COALESCE($4, previous_id),
    hash        = COALESCE($5, hash),
    transaction = COALESCE($6, transaction)
WHERE id = $1;