package service

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/adapters"
	"github.com/alejandroik/trazavino-api/internal/app"
)

// TODO implement
func NewApplication(ctx context.Context) app.Application {
	db, err := adapters.NewMysqlConnection()
	if err != nil {
		panic(err)
	}

	_ = adapters.NewReceptionMysqlRepository(db)
	_ = adapters.NewMacerationMysqlRepository(db)

	return app.Application{}
}
