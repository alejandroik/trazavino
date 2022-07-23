// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: truck.sql

package generated

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addTruck = `-- name: AddTruck :exec
INSERT INTO truck (id, created_at, name)
VALUES ($1, $2, $3)
`

type AddTruckParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Name      string
}

func (q *Queries) AddTruck(ctx context.Context, arg AddTruckParams) error {
	_, err := q.db.ExecContext(ctx, addTruck, arg.ID, arg.CreatedAt, arg.Name)
	return err
}

const getTruck = `-- name: GetTruck :one
SELECT id, created_at, updated_at, deleted_at, name
FROM truck
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetTruck(ctx context.Context, id uuid.UUID) (Truck, error) {
	row := q.db.QueryRowContext(ctx, getTruck, id)
	var i Truck
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
	)
	return i, err
}

const listTrucks = `-- name: ListTrucks :many
SELECT id, created_at, updated_at, deleted_at, name
FROM truck
ORDER BY created_at DESC
OFFSET $1 LIMIT $2
`

type ListTrucksParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) ListTrucks(ctx context.Context, arg ListTrucksParams) ([]Truck, error) {
	rows, err := q.db.QueryContext(ctx, listTrucks, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Truck
	for rows.Next() {
		var i Truck
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

const updateTruck = `-- name: UpdateTruck :exec
UPDATE truck
SET name       = $2,
    updated_at = $3
WHERE id = $1
`

type UpdateTruckParams struct {
	ID        uuid.UUID
	Name      string
	UpdatedAt sql.NullTime
}

func (q *Queries) UpdateTruck(ctx context.Context, arg UpdateTruckParams) error {
	_, err := q.db.ExecContext(ctx, updateTruck, arg.ID, arg.Name, arg.UpdatedAt)
	return err
}
