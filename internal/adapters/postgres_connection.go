package adapters

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func NewPostgresConnection(ctx context.Context) (*pgx.Conn, error) {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		return nil, errors.New("empty db url")
	}

	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to db")
	}

	return db, nil
}
