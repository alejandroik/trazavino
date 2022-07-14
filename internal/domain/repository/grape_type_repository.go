package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type GrapeTypeRepository interface {
	AddGrapeType(ctx context.Context, grapeType *entity.GrapeType) error
	GetGrapeType(ctx context.Context, grapeTypeId int64) (*entity.GrapeType, error)
}
