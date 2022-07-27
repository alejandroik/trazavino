-- name: GetVineyard :one
SELECT *
FROM vineyard
WHERE id = $1
LIMIT 1;

-- name: ListVineyards :many
SELECT *
FROM vineyard
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddVineyard :exec
INSERT INTO vineyard (id, created_at, winery_id, name)
VALUES ($1, $2, $3, $4);

-- name: UpdateVineyard :exec
UPDATE vineyard
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at)
WHERE id = $1;