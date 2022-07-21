package process

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	ProcessId uint
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("process '%d' not found", e.ProcessId)
}

type Repository interface {
	GetProcess(ctx context.Context, processId uint) (*Process, error)
	GetAllProcesses() ([]*Process, error)
}
