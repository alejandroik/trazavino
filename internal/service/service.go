package service

import (
	"context"

	"github.com/alejandroik/trazavino/internal/adapters/dynamodb"
	"github.com/alejandroik/trazavino/internal/app"
	"github.com/alejandroik/trazavino/internal/app/usecase"
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
	macerationRepository := dynamodb.NewMacerationDynamodbRepository(dbClient)

	return app.Application{
		UseCases: app.UseCases{
			RegisterReception:  usecase.NewRegisterReceptionHandler(receptionRepository, log),
			RegisterMaceration: usecase.NewRegisterMacerationHandler(macerationRepository, log),
		},
	}
}
