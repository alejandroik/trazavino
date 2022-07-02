package reception

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	ReceptionId int64
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("reception '%d' not found", e.ReceptionId)
}

type Repository interface {
	AddReception(ctx context.Context, rc *Reception) error
	GetReception(ctx context.Context, receptionId int64) (*Reception, error)
	GetAllReceptions() ([]*Reception, error)
	UpdateReception(ctx context.Context, receptionId int64, updateFn func(ctx context.Context, rc *Reception) (*Reception, error))
}
