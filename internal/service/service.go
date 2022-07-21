package service

import (
	"context"

	"github.com/alejandroik/trazavino/internal/adapters/dynamodb"
	"github.com/alejandroik/trazavino/internal/app"
	"github.com/alejandroik/trazavino/internal/app/command"
	"github.com/alejandroik/trazavino/internal/app/query"
	"github.com/alejandroik/trazavino/pkg/logger"
)

func NewApplication(ctx context.Context, log logger.Interface) app.Application {
	return newApplication(ctx, log)
}

func newApplication(ctx context.Context, log logger.Interface) app.Application {
	dbClient, err := dynamodb.NewDynamoDbClient(ctx)
	if err != nil {
		panic(err)
	}

	receptionRepository := dynamodb.NewReceptionDynamoDbRepository(dbClient)

	return app.Application{
		Commands: app.Commands{
			RegisterReception: command.NewRegisterReceptionHandler(receptionRepository, log),
		},
		Queries: app.Queries{
			ReceptionByID: query.NewReceptionByIDHandler(receptionRepository),
		},
	}
}
