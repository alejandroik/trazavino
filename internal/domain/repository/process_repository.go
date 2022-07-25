package repository

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
)

type ProcessRepository interface {
	GetProcess(ctx context.Context, uuid string) (*entity.Process, error)
	ListProcesses(ctx context.Context, offset int32, limit int32) ([]*entity.Process, error)
}
