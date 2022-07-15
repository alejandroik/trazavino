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

-- name: UpdateProcessHashTransaction :exec
UPDATE process
SET hash        = $2,
    transaction = $3
WHERE id = $1;

-- name: UpdateProcessEndDatePreviousID :exec
UPDATE process
SET end_date    = $2,
    previous_id = $3
WHERE id = $1;