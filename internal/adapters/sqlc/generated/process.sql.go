// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: process.sql

package generated

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addProcess = `-- name: AddProcess :exec
INSERT INTO process (id, created_at, winery_id, start_time, p_type, previous_id)
VALUES ($1, $2, $3, $4, $5, $6)
`

type AddProcessParams struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	WineryID   uuid.UUID
	StartTime  time.Time
	PType      string
	PreviousID uuid.NullUUID
}

func (q *Queries) AddProcess(ctx context.Context, arg AddProcessParams) error {
	_, err := q.db.ExecContext(ctx, addProcess,
		arg.ID,
		arg.CreatedAt,
		arg.WineryID,
		arg.StartTime,
		arg.PType,
		arg.PreviousID,
	)
	return err
}

const getProcess = `-- name: GetProcess :one
SELECT id, created_at, updated_at, deleted_at, winery_id, start_time, end_time, hash, p_type, transaction, previous_id
FROM process
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetProcess(ctx context.Context, id uuid.UUID) (Process, error) {
	row := q.db.QueryRowContext(ctx, getProcess, id)
	var i Process
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.WineryID,
		&i.StartTime,
		&i.EndTime,
		&i.Hash,
		&i.PType,
		&i.Transaction,
		&i.PreviousID,
	)
	return i, err
}

const listProcesses = `-- name: ListProcesses :many
SELECT id, created_at, updated_at, deleted_at, winery_id, start_time, end_time, hash, p_type, transaction, previous_id
FROM process
ORDER BY created_at DESC
OFFSET $1 LIMIT $2
`

type ListProcessesParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) ListProcesses(ctx context.Context, arg ListProcessesParams) ([]Process, error) {
	rows, err := q.db.QueryContext(ctx, listProcesses, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Process
	for rows.Next() {
		var i Process
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.WineryID,
			&i.StartTime,
			&i.EndTime,
			&i.Hash,
			&i.PType,
			&i.Transaction,
			&i.PreviousID,
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

const updateProcess = `-- name: UpdateProcess :exec
UPDATE process
SET updated_at  = COALESCE($2, updated_at),
    end_time    = COALESCE($3, end_time),
    previous_id = COALESCE($4, previous_id),
    hash        = COALESCE($5, hash),
    transaction = COALESCE($6, transaction)
WHERE id = $1
`

type UpdateProcessParams struct {
	ID          uuid.UUID
	UpdatedAt   sql.NullTime
	EndTime     sql.NullTime
	PreviousID  uuid.NullUUID
	Hash        sql.NullString
	Transaction sql.NullString
}

func (q *Queries) UpdateProcess(ctx context.Context, arg UpdateProcessParams) error {
	_, err := q.db.ExecContext(ctx, updateProcess,
		arg.ID,
		arg.UpdatedAt,
		arg.EndTime,
		arg.PreviousID,
		arg.Hash,
		arg.Transaction,
	)
	return err
}
