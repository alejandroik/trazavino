package adapters

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/reception"
)

type ReceptionMysqlRepository struct {
}

func (r ReceptionMysqlRepository) AddReception(ctx context.Context, rc *reception.Reception) error {
	return nil
}

func (r ReceptionMysqlRepository) GetReception(ctx context.Context, receptionId int64) (*reception.Reception, error) {
	return nil, nil
}

func (r ReceptionMysqlRepository) UpdateReception(ctx context.Context, receptionId int64, updateFn func(ctx context.Context, rc *reception.Reception) (*reception.Reception, error)) {
}
