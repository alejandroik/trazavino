package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
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

func (r ProcessRepository) ListProcesses(ctx context.Context, offset int32, limit int32) ([]*entity.Process, error) {
	q := generated.New(r.db)
	pms, err := q.ListProcesses(ctx, generated.ListProcessesParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	var processes []*entity.Process
	for _, pm := range pms {
		process, err := unmarshalProcess(pm)
		if err != nil {
			return nil, err
		}

		processes = append(processes, process)
	}

	return processes, nil
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
		tx.Rollback()
		return nil, err
	}

	return insertedProcess, tx.Commit()
}

func (r ProcessRepository) UpdateProcess(
	ctx context.Context,
	processId int64,
	updateFn func(ctx context.Context, process *entity.Process) (*entity.Process, error),
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := generated.New(tx)

	pm, err := q.GetProcess(ctx, processId)
	if err != nil {
		tx.Rollback()
		return err
	}

	process, err := unmarshalProcess(pm)
	if err != nil {
		tx.Rollback()
		return err
	}

	updatedProcess, err := updateFn(ctx, process)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = q.UpdateProcess(ctx, marshalProcessUpdateParams(updatedProcess))
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func marshalProcessUpdateParams(pr *entity.Process) generated.UpdateProcessParams {
	var endDate sql.NullTime
	if !pr.EndDate().IsZero() {
		endDate = sql.NullTime{Time: pr.EndDate(), Valid: true}
	}

	var previousID sql.NullInt64
	if pr.PreviousUUID() != 0 {
		previousID = sql.NullInt64{Int64: pr.PreviousUUID(), Valid: true}
	}

	var hash sql.NullString
	if pr.Hash() != "" {
		hash = sql.NullString{String: pr.Hash(), Valid: true}
	}

	var transaction sql.NullString
	if pr.Transaction() != "" {
		transaction = sql.NullString{String: pr.Transaction(), Valid: true}
	}

	return generated.UpdateProcessParams{
		ID:          pr.UUID(),
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		EndDate:     endDate,
		PreviousID:  previousID,
		Hash:        hash,
		Transaction: transaction,
	}
}

func unmarshalProcess(pm generated.Process) (*entity.Process, error) {
	return entity.UnmarshalProcessFromDatabase(pm.ID, pm.StartDate, pm.EndDate.Time, pm.Hash.String, pm.Transaction.String, pm.PType, pm.PreviousID.Int64)
}
