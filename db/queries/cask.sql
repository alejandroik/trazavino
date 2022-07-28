-- name: GetCask :one
SELECT *
FROM cask
WHERE id = $1
LIMIT 1;

-- name: ListCasks :many
SELECT *
FROM cask
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: AddCask :exec
INSERT INTO cask (id, created_at, winery_id, name, c_type, is_empty)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateCaskUsage :exec
UPDATE cask
SET is_empty   = $2,
    updated_at = $3
WHERE id = $1;

-- name: UpdateCaskData :exec
UPDATE cask
SET name       = $2,
    c_type     = $3,
    updated_at = $4
WHERE id = $1;