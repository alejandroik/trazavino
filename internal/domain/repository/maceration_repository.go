package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type MacerationRepository interface {
	AddMaceration(ctx context.Context, m *entity.Maceration) error
	GetMaceration(ctx context.Context, macerationId int64) (*entity.Maceration, error)
	UpdateMaceration(ctx context.Context, macerationId int64, updateFn func(ctx context.Context, m *entity.Maceration) (*entity.Maceration, error))
}
