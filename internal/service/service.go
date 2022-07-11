package service

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc"
	"github.com/alejandroik/trazavino-api/internal/app"
)

// TODO implement
func NewApplication(ctx context.Context) (app.Application, func()) {
	return newApplication(ctx), func() {}
}

func newApplication(ctx context.Context) app.Application {
	db, err := sqlc.NewMysqlConnection()
	if err != nil {
		panic(err)
	}

	_ = sqlc.NewProcessMysqlRepository(db)
	//receptionRepository := sqlc.NewReceptionMysqlRepository(db)
	//_ = sqlc.NewMacerationMysqlRepository(db)

	return app.Application{
		Commands: app.Commands{
			//RegisterReception: command.NewRegisterReceptionHandler(receptionRepository),
		},
	}
}
