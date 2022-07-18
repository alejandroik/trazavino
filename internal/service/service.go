package service

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/adapters/dynamodb"
	"github.com/alejandroik/trazavino-api/internal/app"
	"github.com/alejandroik/trazavino-api/internal/app/command"
)

// TODO implement
func NewApplication(ctx context.Context) (app.Application, func()) {
	return newApplication(ctx), func() {}
}

func newApplication(ctx context.Context) app.Application {
	client, err := dynamodb.NewDynamoDbClient(ctx)
	if err != nil {
		panic(err)
	}

	receptionRepository := dynamodb.NewReceptionDynamoDbRepository(client)

	return app.Application{
		Commands: app.Commands{
			RegisterReception: command.NewRegisterReceptionHandler(receptionRepository),
			//RegisterMaceration: command.NewRegisterMacerationHandler(processRepository, macerationRepository, warehouseRepository),
		},
	}
}
