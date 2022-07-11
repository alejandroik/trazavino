package service

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/adapters/gorm"
	"github.com/alejandroik/trazavino-api/internal/app"
	"github.com/alejandroik/trazavino-api/internal/app/command"
)

// TODO implement
func NewApplication(ctx context.Context) (app.Application, func()) {
	return newApplication(ctx), func() {}
}

func newApplication(ctx context.Context) app.Application {
	db, err := gorm.NewMysqlConnection()
	if err != nil {
		panic(err)
	}

	_ = gorm.NewProcessMysqlRepository(db)
	receptionRepository := gorm.NewReceptionMysqlRepository(db)
	_ = gorm.NewMacerationMysqlRepository(db)

	return app.Application{
		Commands: app.Commands{
			RegisterReception: command.NewRegisterReceptionHandler(receptionRepository),
		},
	}
}
