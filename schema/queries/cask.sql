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
INSERT INTO cask (id, created_at, name, c_type, is_empty)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateCask :exec
UPDATE cask
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at),
    c_type     = COALESCE($4, c_type),
    is_empty   = COALESCE($5, is_empty)
WHERE id = $1;