package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type ReceptionRepository interface {
	AddReception(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)
	GetReception(ctx context.Context, receptionId int) (*entity.Reception, error)
	GetAllReceptions() ([]*entity.Reception, error)
	UpdateReception(ctx context.Context, receptionId int, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error
}
