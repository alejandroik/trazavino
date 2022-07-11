package sqlc

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ProcessMysqlRepository struct {
	queries *generated.Queries
}

func NewProcessMysqlRepository(db *sqlx.DB) *ProcessMysqlRepository {
	if db == nil {
		panic("missing db")
	}

	return &ProcessMysqlRepository{queries: generated.New(db)}
}

func (r ProcessMysqlRepository) GetProcess(ctx context.Context, id uint) (*entity.Process, error) {
	pm, err := r.queries.GetProcess(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	process, err := unmarshalProcess(pm)
	if err != nil {
		return nil, err
	}

	return process, nil
}

func (r ProcessMysqlRepository) GetAllProcesses() ([]*entity.Process, error) {
	return nil, nil
}

func unmarshalProcess(pm generated.Process) (*entity.Process, error) {
	return entity.UnmarshalProcessFromDatabase(int(pm.ID), &pm.StartDate, &pm.EndDate, pm.PType, pm.Hash, pm.Transaction, int(pm.PreviousID.Int64))
}
