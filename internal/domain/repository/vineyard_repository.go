package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type VineyardRepository interface {
	AddVineyard(ctx context.Context, vineyard *entity.Vineyard) (*entity.Vineyard, error)
	GetVineyard(ctx context.Context, vineyardId int64) (*entity.Vineyard, error)
}
