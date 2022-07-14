package service

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc"
	"github.com/alejandroik/trazavino-api/internal/app"
	"github.com/alejandroik/trazavino-api/internal/app/command"
)

// TODO implement
func NewApplication(ctx context.Context) (app.Application, func()) {
	return newApplication(ctx), func() {}
}

func newApplication(ctx context.Context) app.Application {
	db, err := sqlc.NewPostgresConnection()
	if err != nil {
		panic(err)
	}

	processRepository := sqlc.NewProcessRepository(db)
	receptionRepository := sqlc.NewReceptionRepository(db)

	return app.Application{
		Commands: app.Commands{
			RegisterReception: command.NewRegisterReceptionHandler(processRepository, receptionRepository),
		},
	}
}
