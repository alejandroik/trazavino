package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type ProcessRepository interface {
	GetProcess(ctx context.Context, uuid string) (*entity.Process, error)
	ListProcesses(ctx context.Context, offset int32, limit int32) ([]*entity.Process, error)
	AddProcess(ctx context.Context, process *entity.Process) (*entity.Process, error)
	UpdateProcess(ctx context.Context, uuid string, updateFn func(ctx context.Context, process *entity.Process) (*entity.Process, error)) error
}
