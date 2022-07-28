-- name: GetWine :one
SELECT *
FROM wine
WHERE id = $1
LIMIT 1;

-- name: ListWines :many
SELECT *
FROM wine
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddWine :exec
INSERT INTO wine (id, created_at, winery_id, name)
VALUES ($1, $2, $3, $4);

-- name: UpdateWineData :exec
UPDATE wine
SET name       = $2,
    updated_at = $3
WHERE id = $1;