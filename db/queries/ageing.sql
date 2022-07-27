-- name: GetAgeing :one
SELECT *
FROM ageing
WHERE id = $1
LIMIT 1;

-- name: FindAgeing :one
SELECT ageing.*
FROM ageing
         INNER JOIN process p on ageing.id = p.id
WHERE end_time IS NULL
  AND ageing.cask_id = $1
LIMIT 1;

-- name: ListAgeings :many
SELECT *
FROM ageing
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddAgeing :exec
INSERT INTO ageing (id, created_at, tank_id, cask_id)
VALUES ($1, $2, $3, $4);