package maceration

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	MacerationId int64
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("maceration '%d' not found", e.MacerationId)
}

type Repository interface {
	AddMaceration(ctx context.Context, m *Maceration) error
	GetMaceration(ctx context.Context, macerationId int64) (*Maceration, error)
	UpdateMaceration(ctx context.Context, macerationId int64, updateFn func(ctx context.Context, m *Maceration) (*Maceration, error))
}
