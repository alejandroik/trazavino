// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: reception.sql

package generated

import (
	"context"
	"time"
)

const addReception = `-- name: AddReception :one
INSERT INTO reception (id, created_at, weight, sugar, truck_id, vineyard_id, grape_type_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at, deleted_at, weight, sugar, truck_id, vineyard_id, grape_type_id
`

type AddReceptionParams struct {
	ID          int64
	CreatedAt   time.Time
	Weight      int32
	Sugar       int32
	TruckID     int64
	VineyardID  int64
	GrapeTypeID int64
}

func (q *Queries) AddReception(ctx context.Context, arg AddReceptionParams) (Reception, error) {
	row := q.db.QueryRowContext(ctx, addReception,
		arg.ID,
		arg.CreatedAt,
		arg.Weight,
		arg.Sugar,
		arg.TruckID,
		arg.VineyardID,
		arg.GrapeTypeID,
	)
	var i Reception
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Weight,
		&i.Sugar,
		&i.TruckID,
		&i.VineyardID,
		&i.GrapeTypeID,
	)
	return i, err
}

const getReception = `-- name: GetReception :one
SELECT id, created_at, updated_at, deleted_at, weight, sugar, truck_id, vineyard_id, grape_type_id
FROM reception
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetReception(ctx context.Context, id int64) (Reception, error) {
	row := q.db.QueryRowContext(ctx, getReception, id)
	var i Reception
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Weight,
		&i.Sugar,
		&i.TruckID,
		&i.VineyardID,
		&i.GrapeTypeID,
	)
	return i, err
}
