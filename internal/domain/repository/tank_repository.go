package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type TankRepository interface {
	AddTank(ctx context.Context, tank *entity.Tank) error
	GetTank(ctx context.Context, tankId string) (*entity.Tank, error)
	ListTanks(ctx context.Context, offset int32, limit int32) ([]*entity.Tank, error)
	UpdateTank(
		ctx context.Context,
		tankId string,
		updateFn func(ctx context.Context, wh *entity.Tank) (*entity.Tank, error),
	) error
}
