package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type BottlingRepository interface {
	AddBottling(ctx context.Context, b *entity.Bottling) error
	GetBottling(ctx context.Context, bottlingUUID string) (*entity.Bottling, error)
	UpdateBottling(
		ctx context.Context,
		bottlingUUID string,
		updateFn func(ctx context.Context, b *entity.Bottling) (*entity.Bottling, error),
	) error
}
