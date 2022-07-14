package sqlc

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type TruckRepository struct {
	db *sqlx.DB
}

func NewTruckRepository(db *sqlx.DB) *TruckRepository {
	if db == nil {
		panic("missing db")
	}

	return &TruckRepository{db: db}
}

func (r TruckRepository) AddTruck(ctx context.Context, truck *entity.Truck) (*entity.Truck, error) {
	return nil, nil
}

func (r TruckRepository) GetTruck(ctx context.Context, truckId int64) (*entity.Truck, error) {
	return nil, nil
}
