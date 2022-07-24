package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type BottleRepository interface {
	AddBottle(ctx context.Context, bottle *entity.Bottle) error
	GetBottle(ctx context.Context, bottleId string) (*entity.Bottle, error)
	ListBottles(ctx context.Context, offset int32, limit int32) ([]*entity.Bottle, error)
	UpdateBottle(
		ctx context.Context,
		bottleId string,
		updateFn func(ctx context.Context, b *entity.Bottle) (*entity.Bottle, error),
	) error
}
