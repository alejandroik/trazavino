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

-- name: UpdateVineyardData :exec
UPDATE vineyard
SET name       = $2,
    updated_at = $3
WHERE id = $1;