package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type FermentationRepository interface {
	AddFermentation(ctx context.Context, f *entity.Fermentation) error
	GetFermentation(ctx context.Context, fermentationUUID string) (*entity.Fermentation, error)
	UpdateFermentation(
		ctx context.Context,
		fermentationUUID string,
		updateFn func(ctx context.Context, f *entity.Fermentation) (*entity.Fermentation, error),
	) error
}
