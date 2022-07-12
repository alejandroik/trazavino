package sqlc

import (
	"context"
	"database/sql"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ProcessMysqlRepository struct {
	db *sqlx.DB
}

func NewProcessMysqlRepository(db *sqlx.DB) *ProcessMysqlRepository {
	if db == nil {
		panic("missing db")
	}

	return &ProcessMysqlRepository{db: db}
}

func (r ProcessMysqlRepository) GetProcess(ctx context.Context, id uint) (*entity.Process, error) {
	q := generated.New(r.db)
	pm, err := q.GetProcess(ctx, int64(id))
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

func (r ProcessMysqlRepository) AddProcess(ctx context.Context, process *entity.Process) (*entity.Process, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	q := generated.New(tx)

	result, err := q.AddProcess(ctx, generated.AddProcessParams{
		StartDate: process.StartDate(),
		PType:     process.Ptype(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var hash sql.NullString
	if process.Hash() != "" {
		hash = sql.NullString{
			String: process.Hash(),
			Valid:  true,
		}
	}

	var previousID sql.NullInt64
	if process.PreviousId() > 0 {
		previousID = sql.NullInt64{
			Int64: int64(process.PreviousId()),
			Valid: true,
		}
	}

	pm := generated.Process{
		ID:          insertedId,
		StartDate:   process.StartDate(),
		EndDate:     process.EndDate(),
		Hash:        hash,
		PType:       process.Ptype(),
		Transaction: process.Transaction(),
		PreviousID:  previousID,
	}

	insertedProcess, err := unmarshalProcess(pm)
	if err != nil {
		return nil, err
	}

	return insertedProcess, tx.Commit()
}

func unmarshalProcess(pm generated.Process) (*entity.Process, error) {
	return entity.UnmarshalProcessFromDatabase(int(pm.ID), &pm.StartDate, &pm.EndDate, pm.PType, pm.Hash.String, pm.Transaction, int(pm.PreviousID.Int64))
}
