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

-- name: UpdateBottleData :exec
UPDATE bottle
SET name       = $2,
    updated_at = $3
WHERE id = $1;