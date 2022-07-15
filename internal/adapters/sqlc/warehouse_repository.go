package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type WarehouseRepository struct {
	db *sqlx.DB
}

func NewWarehouseRepository(db *sqlx.DB) *WarehouseRepository {
	if db == nil {
		panic("missing db")
	}

	return &WarehouseRepository{db: db}
}

func (r WarehouseRepository) AddWarehouse(ctx context.Context, warehouse *entity.Warehouse) (*entity.Warehouse, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	q := generated.New(tx)

	insertedId, err := q.AddWarehouse(ctx, generated.AddWarehouseParams{
		CreatedAt: time.Now(),
		Name:      warehouse.Name(),
		IsEmpty:   warehouse.IsEmpty(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	wm := generated.Warehouse{
		ID:      insertedId,
		Name:    warehouse.Name(),
		IsEmpty: warehouse.IsEmpty(),
	}

	insertedWarehouse, err := unmarshalWarehouse(wm)
	if err != nil {
		return nil, err
	}

	return insertedWarehouse, tx.Commit()
}

func (r WarehouseRepository) GetWarehouse(ctx context.Context, warehouseId int64) (*entity.Warehouse, error) {
	q := generated.New(r.db)
	wm, err := q.GetWarehouse(ctx, warehouseId)
	if err != nil {
		return nil, err
	}

	warehouse, err := unmarshalWarehouse(wm)
	if err != nil {
		return nil, err
	}

	return warehouse, nil
}

func unmarshalWarehouse(wm generated.Warehouse) (*entity.Warehouse, error) {
	return entity.UnmarshalWarehouseFromDatabase(wm.ID, wm.Name, wm.IsEmpty)
}
