package sqlc

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
)

type TruckRepository struct {
	db *pgx.Conn
}

func NewTruckRepository(db *pgx.Conn) *TruckRepository {
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
	wineryUuid, err := uuid.Parse(truck.WineryUUID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddTruck(ctx, generated.AddTruckParams{
		ID:        truckUuid,
		CreatedAt: time.Now(),
		Name:      truck.Name(),
		WineryID:  wineryUuid,
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
