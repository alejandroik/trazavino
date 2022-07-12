// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: reception.sql

package generated

import (
	"context"
	"database/sql"
)

const addReception = `-- name: AddReception :execresult
INSERT INTO reception (id, weight, sugar)
VALUES (?, ?, ?)
`

type AddReceptionParams struct {
	ID     int64
	Weight int32
	Sugar  int32
}

func (q *Queries) AddReception(ctx context.Context, arg AddReceptionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addReception, arg.ID, arg.Weight, arg.Sugar)
}

const getReception = `-- name: GetReception :one
SELECT id, weight, sugar
FROM reception
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetReception(ctx context.Context, id int64) (Reception, error) {
	row := q.db.QueryRowContext(ctx, getReception, id)
	var i Reception
	err := row.Scan(&i.ID, &i.Weight, &i.Sugar)
	return i, err
}