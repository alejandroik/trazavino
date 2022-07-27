package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type AgeingRepository interface {
	AddAgeing(ctx context.Context, f *entity.Ageing) error
	GetAgeing(ctx context.Context, ageingUUID string) (*entity.Ageing, error)
	UpdateAgeing(
		ctx context.Context,
		ageingUUID string,
		updateFn func(ctx context.Context, f *entity.Ageing) (*entity.Ageing, error),
	) error
}
