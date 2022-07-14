package sqlc

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewPostgresConnection() (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to MySQL")
	}

	return db, nil
}
