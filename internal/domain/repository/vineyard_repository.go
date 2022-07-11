package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type VineyardReception interface {
	AddVineyard(ctx context.Context, vineyard *entity.Vineyard) error
	GetVineyard(ctx context.Context, vineyardId int) (*entity.Vineyard, error)
}
