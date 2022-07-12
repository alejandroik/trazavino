// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: process.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const addProcess = `-- name: AddProcess :execresult
INSERT INTO process (start_date, p_type)
VALUES (?, ?)
`

type AddProcessParams struct {
	StartDate time.Time
	PType     string
}

func (q *Queries) AddProcess(ctx context.Context, arg AddProcessParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addProcess, arg.StartDate, arg.PType)
}

const getProcess = `-- name: GetProcess :one
SELECT id, start_date, end_date, hash, p_type, transaction, previous_id
FROM process
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetProcess(ctx context.Context, id int64) (Process, error) {
	row := q.db.QueryRowContext(ctx, getProcess, id)
	var i Process
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.Hash,
		&i.PType,
		&i.Transaction,
		&i.PreviousID,
	)
	return i, err
}