// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package generated

import (
	"database/sql"
	"time"
)

type Process struct {
	ID          int64
	StartDate   time.Time
	EndDate     time.Time
	Hash        sql.NullString
	PType       string
	Transaction string
	PreviousID  sql.NullInt64
}

type Reception struct {
	ID     int64
	Weight int32
	Sugar  int32
}
