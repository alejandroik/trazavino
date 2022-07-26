// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: winery.sql

package generated

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addWinery = `-- name: AddWinery :exec
INSERT INTO winery (id, created_at, name)
VALUES ($1, $2, $3)
`

type AddWineryParams struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Name      string    `db:"name"`
}

func (q *Queries) AddWinery(ctx context.Context, arg AddWineryParams) error {
	_, err := q.db.Exec(ctx, addWinery, arg.ID, arg.CreatedAt, arg.Name)
	return err
}

const getWinery = `-- name: GetWinery :one
SELECT id, created_at, updated_at, deleted_at, name
FROM winery
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetWinery(ctx context.Context, id uuid.UUID) (Winery, error) {
	row := q.db.QueryRow(ctx, getWinery, id)
	var i Winery
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
	)
	return i, err
}

const listWinerys = `-- name: ListWinerys :many
SELECT id, created_at, updated_at, deleted_at, name
FROM winery
ORDER BY created_at DESC
    OFFSET $1 LIMIT $2
`

type ListWinerysParams struct {
	Offset int32 `db:"offset"`
	Limit  int32 `db:"limit"`
}

func (q *Queries) ListWinerys(ctx context.Context, arg ListWinerysParams) ([]Winery, error) {
	rows, err := q.db.Query(ctx, listWinerys, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Winery
	for rows.Next() {
		var i Winery
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateWinery = `-- name: UpdateWinery :exec
UPDATE winery
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at)
WHERE id = $1
`

type UpdateWineryParams struct {
	ID        uuid.UUID    `db:"id"`
	Name      string       `db:"name"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (q *Queries) UpdateWinery(ctx context.Context, arg UpdateWineryParams) error {
	_, err := q.db.Exec(ctx, updateWinery, arg.ID, arg.Name, arg.UpdatedAt)
	return err
}
