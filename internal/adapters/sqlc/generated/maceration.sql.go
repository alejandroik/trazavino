// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: maceration.sql

package generated

import (
	"context"
	"time"
)

const addMaceration = `-- name: AddMaceration :one
INSERT INTO maceration (id, created_at, reception_id, warehouse_id)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, deleted_at, reception_id, warehouse_id
`

type AddMacerationParams struct {
	ID          int64
	CreatedAt   time.Time
	ReceptionID int64
	WarehouseID int64
}

func (q *Queries) AddMaceration(ctx context.Context, arg AddMacerationParams) (Maceration, error) {
	row := q.db.QueryRowContext(ctx, addMaceration,
		arg.ID,
		arg.CreatedAt,
		arg.ReceptionID,
		arg.WarehouseID,
	)
	var i Maceration
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ReceptionID,
		&i.WarehouseID,
	)
	return i, err
}

const getMaceration = `-- name: GetMaceration :one
SELECT id, created_at, updated_at, deleted_at, reception_id, warehouse_id
FROM maceration
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetMaceration(ctx context.Context, id int64) (Maceration, error) {
	row := q.db.QueryRowContext(ctx, getMaceration, id)
	var i Maceration
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ReceptionID,
		&i.WarehouseID,
	)
	return i, err
}

const listMacerations = `-- name: ListMacerations :many
SELECT id, created_at, updated_at, deleted_at, reception_id, warehouse_id
FROM maceration
ORDER BY id
OFFSET $1 LIMIT $2
`

type ListMacerationsParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) ListMacerations(ctx context.Context, arg ListMacerationsParams) ([]Maceration, error) {
	rows, err := q.db.QueryContext(ctx, listMacerations, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Maceration
	for rows.Next() {
		var i Maceration
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ReceptionID,
			&i.WarehouseID,
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
