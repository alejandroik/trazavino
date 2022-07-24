package repository

import (
	"context"
	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type WineryRepository interface {
	AddWinery(ctx context.Context, winery *entity.Winery) error
	GetWinery(ctx context.Context, wineryId string) (*entity.Winery, error)
	LIstWineries(ctx context.Context, offset int32, limit int32) ([]*entity.Winery, error)
	UpdateWinery(
		ctx context.Context,
		wineryId string,
		updateFn func(ctx context.Context, w *entity.Winery) (*entity.Winery, error),
	) error
}
