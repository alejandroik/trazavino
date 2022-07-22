// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: vineyard.sql

package generated

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addVineyard = `-- name: AddVineyard :exec
INSERT INTO vineyard (id, created_at, name)
VALUES ($1, $2, $3)
`

type AddVineyardParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Name      string
}

func (q *Queries) AddVineyard(ctx context.Context, arg AddVineyardParams) error {
	_, err := q.db.ExecContext(ctx, addVineyard, arg.ID, arg.CreatedAt, arg.Name)
	return err
}

const getVineyard = `-- name: GetVineyard :one
SELECT id, created_at, updated_at, deleted_at, name
FROM vineyard
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetVineyard(ctx context.Context, id uuid.UUID) (Vineyard, error) {
	row := q.db.QueryRowContext(ctx, getVineyard, id)
	var i Vineyard
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
	)
	return i, err
}

const listVineyards = `-- name: ListVineyards :many
SELECT id, created_at, updated_at, deleted_at, name
FROM vineyard
ORDER BY created_at DESC
OFFSET $1 LIMIT $2
`

type ListVineyardsParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) ListVineyards(ctx context.Context, arg ListVineyardsParams) ([]Vineyard, error) {
	rows, err := q.db.QueryContext(ctx, listVineyards, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Vineyard
	for rows.Next() {
		var i Vineyard
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
