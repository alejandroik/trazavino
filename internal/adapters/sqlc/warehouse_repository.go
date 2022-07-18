package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
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

	wm, err := q.AddWarehouse(ctx, generated.AddWarehouseParams{
		CreatedAt: time.Now(),
		Name:      warehouse.Name(),
		IsEmpty:   true,
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	insertedWarehouse, err := unmarshalWarehouse(wm)
	if err != nil {
		tx.Rollback()
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

func (r WarehouseRepository) ListWarehouses(ctx context.Context, offset int32, limit int32) ([]*entity.Warehouse, error) {
	q := generated.New(r.db)
	wms, err := q.ListWarehouses(ctx, generated.ListWarehousesParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	var warehouses []*entity.Warehouse
	for _, wm := range wms {
		warehouse, err := unmarshalWarehouse(wm)
		if err != nil {
			return nil, err
		}

		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}

func (r WarehouseRepository) UpdateWarehouse(
	ctx context.Context,
	warehouseId int64,
	updateFn func(ctx context.Context, wh *entity.Warehouse) (*entity.Warehouse, error),
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := generated.New(tx)

	wm, err := q.GetWarehouse(ctx, warehouseId)
	if err != nil {
		tx.Rollback()
		return err
	}

	warehouse, err := unmarshalWarehouse(wm)
	if err != nil {
		tx.Rollback()
		return err
	}

	updatedWarehouse, err := updateFn(ctx, warehouse)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = q.UpdateWarehouse(ctx, marshalWarehouseUpdateParams(updatedWarehouse))
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func marshalWarehouseUpdateParams(wh *entity.Warehouse) generated.UpdateWarehouseParams {
	return generated.UpdateWarehouseParams{
		ID:        wh.ID(),
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		IsEmpty:   wh.IsEmpty(),
	}
}

func unmarshalWarehouse(wm generated.Warehouse) (*entity.Warehouse, error) {
	return entity.UnmarshalWarehouseFromDatabase(wm.ID, wm.Name, wm.IsEmpty)
}
