package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type ProcessRepository interface {
	GetProcess(ctx context.Context, processId uint) (*entity.Process, error)
	GetAllProcesses() ([]*entity.Process, error)
}
