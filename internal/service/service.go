package service

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/adapters/mysql"
	"github.com/alejandroik/trazavino-api/internal/app"
	"github.com/alejandroik/trazavino-api/internal/app/command"
)

// TODO implement
func NewApplication(ctx context.Context) (app.Application, func()) {
	return newApplication(ctx), func() {}
}

func newApplication(ctx context.Context) app.Application {
	db, err := mysql.NewMysqlConnection()
	if err != nil {
		panic(err)
	}

	_ = mysql.NewProcessMysqlRepository(db)
	receptionRepository := mysql.NewReceptionMysqlRepository(db)
	_ = mysql.NewMacerationMysqlRepository(db)

	return app.Application{
		Commands: app.Commands{
			RegisterReception: command.NewRegisterReceptionHandler(receptionRepository),
		},
	}
}
