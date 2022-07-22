package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type VineyardRepository interface {
	AddVineyard(ctx context.Context, vineyard *entity.Vineyard) error
	GetVineyard(ctx context.Context, vineyardId string) (*entity.Vineyard, error)
}
