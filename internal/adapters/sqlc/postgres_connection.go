package sqlc

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func NewPostgresConnection(ctx context.Context) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to db")
	}

	return db, nil
}
