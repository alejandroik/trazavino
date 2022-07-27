-- name: GetBottle :one
SELECT *
FROM bottle
WHERE id = $1
LIMIT 1;

-- name: ListBottles :many
SELECT *
FROM bottle
ORDER BY created_at DESC
    OFFSET $1 LIMIT $2;

-- name: AddBottle :exec
INSERT INTO bottle (id, created_at, winery_id, name)
VALUES ($1, $2, $3, $4);

-- name: UpdateBottle :exec
UPDATE bottle
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at)
WHERE id = $1;