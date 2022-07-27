package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type CaskRepository interface {
	AddCask(ctx context.Context, cask *entity.Cask) error
	GetCask(ctx context.Context, caskId string) (*entity.Cask, error)
	ListCasks(ctx context.Context, offset int32, limit int32) ([]*entity.Cask, error)
	UpdateCask(
		ctx context.Context,
		caskId string,
		updateFn func(ctx context.Context, c *entity.Cask) (*entity.Cask, error),
	) error
}
