package migrate

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
)

func RunMigrations() error {
	m, err := migrate.New("file:///"+os.Getenv("PWD")+"/db/migrations", os.Getenv("DB_URL"))
	if err != nil {
		return errors.Wrap(err, "error while running migrations")
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "error while running migrations")
	}

	return nil
}
