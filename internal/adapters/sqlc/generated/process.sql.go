// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: process.sql

package generated

import (
	"context"
	"time"
)

const addProcess = `-- name: AddProcess :one
INSERT INTO process (created_at, start_date, p_type)
VALUES ($1, $2, $3)
RETURNING id, created_at, updated_at, deleted_at, start_date, end_date, hash, p_type, transaction, previous_id
`

type AddProcessParams struct {
	CreatedAt time.Time
	StartDate time.Time
	PType     string
}

func (q *Queries) AddProcess(ctx context.Context, arg AddProcessParams) (Process, error) {
	row := q.db.QueryRowContext(ctx, addProcess, arg.CreatedAt, arg.StartDate, arg.PType)
	var i Process
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.StartDate,
		&i.EndDate,
		&i.Hash,
		&i.PType,
		&i.Transaction,
		&i.PreviousID,
	)
	return i, err
}

const getProcess = `-- name: GetProcess :one
SELECT id, created_at, updated_at, deleted_at, start_date, end_date, hash, p_type, transaction, previous_id
FROM process
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetProcess(ctx context.Context, id int64) (Process, error) {
	row := q.db.QueryRowContext(ctx, getProcess, id)
	var i Process
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.StartDate,
		&i.EndDate,
		&i.Hash,
		&i.PType,
		&i.Transaction,
		&i.PreviousID,
	)
	return i, err
}
