package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type MacerationRepository interface {
	AddMaceration(ctx context.Context, m *entity.Maceration) error
	GetMaceration(ctx context.Context, macerationUUID string) (*entity.Maceration, error)
	UpdateMaceration(ctx context.Context, macerationUUID string, updateFn func(ctx context.Context, m *entity.Maceration) (*entity.Maceration, error)) error
}
