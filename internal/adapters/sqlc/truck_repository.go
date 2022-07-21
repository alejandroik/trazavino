package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
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
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	q := generated.New(tx)

	tm, err := q.AddTruck(ctx, generated.AddTruckParams{
		CreatedAt: time.Now(),
		Name:      truck.Name(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	insertedTruck, err := unmarshalTruck(tm)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return insertedTruck, tx.Commit()
}

func (r TruckRepository) GetTruck(ctx context.Context, truckId int64) (*entity.Truck, error) {
	q := generated.New(r.db)
	tm, err := q.GetTruck(ctx, truckId)
	if err != nil {
		return nil, err
	}

	truck, err := unmarshalTruck(tm)
	if err != nil {
		return nil, err
	}

	return truck, nil
}

func unmarshalTruck(tm generated.Truck) (*entity.Truck, error) {
	return entity.UnmarshalTruckFromDatabase(tm.ID, tm.Name)
}
