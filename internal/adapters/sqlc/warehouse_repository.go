package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
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

func (r WarehouseRepository) AddWarehouse(ctx context.Context, warehouse *entity.Warehouse) error {
	whUuid, err := uuid.Parse(warehouse.ID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddWarehouse(ctx, generated.AddWarehouseParams{
		ID:        whUuid,
		CreatedAt: time.Now(),
		Name:      warehouse.Name(),
		IsEmpty:   true,
	}); err != nil {
		return err
	}

	return nil
}

func (r WarehouseRepository) GetWarehouse(ctx context.Context, warehouseId string) (*entity.Warehouse, error) {
	//q := generated.New(r.db)
	//wm, err := q.GetWarehouse(ctx, warehouseId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//warehouse, err := unmarshalWarehouse(wm)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return warehouse, nil
	return nil, nil
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
	warehouseId string,
	updateFn func(ctx context.Context, wh *entity.Warehouse) (*entity.Warehouse, error),
) error {
	//tx, err := r.db.BeginTx(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//q := generated.New(tx)
	//
	//wm, err := q.GetWarehouse(ctx, warehouseId)
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//warehouse, err := unmarshalWarehouse(wm)
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//updatedWarehouse, err := updateFn(ctx, warehouse)
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//err = q.UpdateWarehouse(ctx, marshalWarehouseUpdateParams(updatedWarehouse))
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//return tx.Commit()
	return nil
}

func marshalWarehouseUpdateParams(wh *entity.Warehouse) generated.UpdateWarehouseParams {
	//return generated.UpdateWarehouseParams{
	//	ID:        wh.ID(),
	//	UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	//	IsEmpty:   wh.IsEmpty(),
	//}
	return generated.UpdateWarehouseParams{}
}

func unmarshalWarehouse(wm generated.Warehouse) (*entity.Warehouse, error) {
	//return entity.UnmarshalWarehouseFromDatabase(wm.ID, wm.Name, wm.IsEmpty)
	return nil, nil
}
