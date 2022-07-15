package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type ReceptionRepository interface {
	AddReception(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)
	GetReception(ctx context.Context, receptionId int64) (*entity.Reception, error)
	ListReceptions(ctx context.Context, offset int32, limit int32) ([]*entity.Reception, error)
	UpdateReception(ctx context.Context, receptionId int64, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error
}
