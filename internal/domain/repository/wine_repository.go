package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type WineRepository interface {
	AddWine(ctx context.Context, wine *entity.Wine) error
	GetWine(ctx context.Context, wineId string) (*entity.Wine, error)
	ListWines(ctx context.Context, offset int32, limit int32) ([]*entity.Wine, error)
	UpdateWine(
		ctx context.Context,
		wineId string,
		updateFn func(ctx context.Context, w *entity.Wine) (*entity.Wine, error),
	) error
}
