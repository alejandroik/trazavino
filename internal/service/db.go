package service

import (
	"context"
	"os"

	"github.com/alejandroik/trazavino/internal/adapters"
	"github.com/alejandroik/trazavino/internal/adapters/dynamodb"
	"github.com/alejandroik/trazavino/internal/adapters/migrate"
	"github.com/alejandroik/trazavino/internal/adapters/sqlc"
	"github.com/alejandroik/trazavino/internal/domain/repository"
)

type repositories struct {
	ReceptionRepository    repository.ReceptionRepository
	MacerationRepository   repository.MacerationRepository
	FermentationRepository repository.FermentationRepository
	AgeingRepository       repository.AgeingRepository
	BottlingRepository     repository.BottlingRepository

	WarehouseRepository repository.WarehouseRepository
	TankRepository      repository.TankRepository
	CaskRepository      repository.CaskRepository
}

func initRepositories(ctx context.Context) (*repositories, error) {
	r := &repositories{}

	dbAdapter := os.Getenv("DB_ADAPTER")
	if dbAdapter == "" {
		dbAdapter = "postgres"
	}

	switch dbAdapter {
	case "postgres":
		db, err := adapters.NewPostgresConnection(ctx)
		if err != nil {
			return nil, err
		}

		if err = migrate.RunMigrations(); err != nil {
			return nil, err
		}

		r.ReceptionRepository = sqlc.NewReceptionRepository(db)
		r.MacerationRepository = sqlc.NewMacerationRepository(db)
		r.FermentationRepository = sqlc.NewFermentationRepository(db)
		r.AgeingRepository = sqlc.NewAgeingRepository(db)
		r.BottlingRepository = sqlc.NewBottlingRepository(db)

		r.WarehouseRepository = sqlc.NewWarehouseRepository(db)
		r.TankRepository = sqlc.NewTankRepository(db)
		r.CaskRepository = sqlc.NewCaskRepository(db)

	case "dynamodb":
		db, err := dynamodb.NewDynamoDbClient(ctx)
		if err != nil {
			return nil, err
		}

		r.ReceptionRepository = dynamodb.NewReceptionDynamoDbRepository(db)
		r.MacerationRepository = dynamodb.NewMacerationDynamodbRepository(db)
	}

	return r, nil
}
