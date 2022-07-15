package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type WarehouseRepository interface {
	AddWarehouse(ctx context.Context, warehouse *entity.Warehouse) (*entity.Warehouse, error)
	GetWarehouse(ctx context.Context, warehouseId int64) (*entity.Warehouse, error)
}
