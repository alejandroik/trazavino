package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ProcessRepository struct {
	db *sqlx.DB
}

func NewProcessRepository(db *sqlx.DB) *ProcessRepository {
	if db == nil {
		panic("missing db")
	}

	return &ProcessRepository{db: db}
}

func (r ProcessRepository) GetProcess(ctx context.Context, id int64) (*entity.Process, error) {
	q := generated.New(r.db)
	pm, err := q.GetProcess(ctx, id)
	if err != nil {
		return nil, err
	}

	process, err := unmarshalProcess(pm)
	if err != nil {
		return nil, err
	}

	return process, nil
}

func (r ProcessRepository) GetAllProcesses() ([]*entity.Process, error) {
	return nil, nil
}

func (r ProcessRepository) AddProcess(ctx context.Context, process *entity.Process) (*entity.Process, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	q := generated.New(tx)

	pm, err := q.AddProcess(ctx, generated.AddProcessParams{
		CreatedAt: time.Now(),
		StartDate: process.StartDate(),
		PType:     process.Ptype(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	insertedProcess, err := unmarshalProcess(pm)
	if err != nil {
		return nil, err
	}

	return insertedProcess, tx.Commit()
}

func unmarshalProcess(pm generated.Process) (*entity.Process, error) {
	return entity.UnmarshalProcessFromDatabase(pm.ID, pm.StartDate, pm.EndDate.Time, pm.PType, pm.Hash.String, pm.Transaction.String, pm.PreviousID.Int64)
}
