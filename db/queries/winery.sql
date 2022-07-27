-- name: GetWinery :one
SELECT *
FROM winery
WHERE id = $1
LIMIT 1;

-- name: ListWinerys :many
SELECT *
FROM winery
ORDER BY created_at DESC
    OFFSET $1 LIMIT $2;

-- name: AddWinery :exec
INSERT INTO winery (id, created_at, name)
VALUES ($1, $2, $3);

-- name: UpdateWinery :exec
UPDATE winery
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at)
WHERE id = $1;