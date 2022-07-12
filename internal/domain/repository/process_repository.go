package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type ProcessRepository interface {
	GetProcess(ctx context.Context, id uint) (*entity.Process, error)
	GetAllProcesses() ([]*entity.Process, error)
	AddProcess(ctx context.Context, process *entity.Process) (*entity.Process, error)
}
