package service

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/adapters/mysql"
	"github.com/alejandroik/trazavino-api/internal/app"
)

// TODO implement
func NewApplication(ctx context.Context) app.Application {
	db, err := mysql.NewMysqlConnection()
	if err != nil {
		panic(err)
	}

	_ = mysql.NewProcessMysqlRepository(db)
	_ = mysql.NewReceptionMysqlRepository(db)
	_ = mysql.NewMacerationMysqlRepository(db)

	return app.Application{}
}
