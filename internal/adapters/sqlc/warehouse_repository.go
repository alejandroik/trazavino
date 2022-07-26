package sqlc

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
)

type WarehouseRepository struct {
	db *pgx.Conn
}

func NewWarehouseRepository(db *pgx.Conn) *WarehouseRepository {
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
	wineryUuid, err := uuid.Parse(warehouse.WineryUUID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddWarehouse(ctx, generated.AddWarehouseParams{
		ID:        whUuid,
		CreatedAt: time.Now(),
		Name:      warehouse.Name(),
		IsEmpty:   true,
		WineryID:  wineryUuid,
	}); err != nil {
		return err
	}

	return nil
}

func (r WarehouseRepository) GetWarehouse(ctx context.Context, warehouseId string) (*entity.Warehouse, error) {
	whUuid, err := uuid.Parse(warehouseId)
	if err != nil {
		return nil, err
	}

	q := generated.New(r.db)
	wm, err := q.GetWarehouse(ctx, whUuid)
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
	warehouseId string,
	updateFn func(ctx context.Context, wh *entity.Warehouse) (*entity.Warehouse, error),
) error {
	whUuid, err := uuid.Parse(warehouseId)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	q := generated.New(tx)

	wm, err := q.GetWarehouse(ctx, whUuid)
	if err != nil {
		return err
	}

	warehouse, err := unmarshalWarehouse(wm)
	if err != nil {
		return err
	}

	updatedWarehouse, err := updateFn(ctx, warehouse)
	if err != nil {
		return err
	}

	if err = q.UpdateWarehouse(ctx, generated.UpdateWarehouseParams{
		ID:        whUuid,
		Name:      updatedWarehouse.Name(),
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		IsEmpty:   updatedWarehouse.IsEmpty(),
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func unmarshalWarehouse(wm generated.Warehouse) (*entity.Warehouse, error) {
	return entity.UnmarshalWarehouseFromDatabase(wm.ID.String(), wm.Name, wm.IsEmpty, wm.WineryID.String())
}
