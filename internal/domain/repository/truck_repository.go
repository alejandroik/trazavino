package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type TruckRepository interface {
	AddTruck(ctx context.Context, truck *entity.Truck) error
	GetTruck(ctx context.Context, truckId int64) (*entity.Truck, error)
}
