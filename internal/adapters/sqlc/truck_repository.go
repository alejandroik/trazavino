package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
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

func (r TruckRepository) AddTruck(ctx context.Context, truck *entity.Truck) error {
	truckUuid, err := uuid.Parse(truck.ID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddTruck(ctx, generated.AddTruckParams{
		ID:        truckUuid,
		CreatedAt: time.Now(),
		Name:      truck.Name(),
	}); err != nil {
		return err
	}

	return nil
}

func (r TruckRepository) GetTruck(ctx context.Context, truckId int64) (*entity.Truck, error) {
	//q := generated.New(r.db)
	//tm, err := q.GetTruck(ctx, truckId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//truck, err := unmarshalTruck(tm)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return truck, nil
	return nil, nil
}

func unmarshalTruck(tm generated.Truck) (*entity.Truck, error) {
	//return entity.UnmarshalTruckFromDatabase(tm.ID, tm.Name)
	return nil, nil
}
