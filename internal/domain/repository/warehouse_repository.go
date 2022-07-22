package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type WarehouseRepository interface {
	AddWarehouse(ctx context.Context, warehouse *entity.Warehouse) error
	GetWarehouse(ctx context.Context, warehouseId int64) (*entity.Warehouse, error)
	ListWarehouses(ctx context.Context, offset int32, limit int32) ([]*entity.Warehouse, error)
	UpdateWarehouse(
		ctx context.Context,
		warehouseId int64,
		updateFn func(ctx context.Context, wh *entity.Warehouse) (*entity.Warehouse, error),
	) error
}
