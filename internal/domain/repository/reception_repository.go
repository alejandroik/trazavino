package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type ReceptionRepository interface {
	AddReception(ctx context.Context, rc *entity.Reception) error
	GetReception(ctx context.Context, receptionUUID string) (*entity.Reception, error)
	UpdateReception(ctx context.Context, receptionUUID string, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error
}
