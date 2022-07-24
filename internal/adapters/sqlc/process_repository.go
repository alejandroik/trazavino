package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
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

func (r ProcessRepository) GetProcess(ctx context.Context, id string) (*entity.Process, error) {
	processUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	q := generated.New(r.db)
	pm, err := q.GetProcess(ctx, processUuid)
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

func (r ProcessRepository) AddProcess(ctx context.Context, pr *entity.Process) error {
	processUuid, err := uuid.Parse(pr.UUID())
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := generated.New(tx)

	if err = q.AddProcess(ctx, generated.AddProcessParams{
		ID:        processUuid,
		CreatedAt: time.Now(),
		StartTime: pr.StartTime(),
		PType:     pr.Ptype(),
	}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r ProcessRepository) UpdateProcess(
	ctx context.Context,
	prUuid string,
	updateFn func(ctx context.Context, process *entity.Process) (*entity.Process, error),
) error {
	processUuid, err := uuid.Parse(prUuid)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := generated.New(tx)

	pm, err := q.GetProcess(ctx, processUuid)
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

	params, err := marshalProcessUpdateParams(updatedProcess)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = q.UpdateProcess(ctx, params)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func marshalProcessUpdateParams(pr *entity.Process) (generated.UpdateProcessParams, error) {
	processUuid, err := uuid.Parse(pr.UUID())
	if err != nil {
		return generated.UpdateProcessParams{}, err
	}

	var endDate sql.NullTime
	if !pr.EndTime().IsZero() {
		endDate = sql.NullTime{Time: pr.EndTime(), Valid: true}
	}

	var previousID uuid.NullUUID
	if pr.PreviousUUID() != "" {
		prevUuid, err := uuid.Parse(pr.PreviousUUID())
		if err != nil {
			return generated.UpdateProcessParams{}, err
		}
		previousID = uuid.NullUUID{UUID: prevUuid, Valid: true}
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
		ID:          processUuid,
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		EndTime:     endDate,
		PreviousID:  previousID,
		Hash:        hash,
		Transaction: transaction,
	}, nil
}

func unmarshalProcess(pm generated.Process) (*entity.Process, error) {
	var endTime time.Time
	if pm.EndTime.Valid {
		endTime = pm.EndTime.Time
	}

	var previousUUID string
	if pm.PreviousID.Valid {
		previousUUID = pm.PreviousID.UUID.String()
	}

	var hash string
	if pm.Hash.Valid {
		hash = pm.Hash.String
	}

	var transaction string
	if pm.Transaction.Valid {
		transaction = pm.Transaction.String
	}

	return entity.UnmarshalProcessFromDatabase(
		pm.ID.String(),
		pm.StartTime,
		pm.PType,
		endTime,
		previousUUID,
		hash,
		transaction,
	)
}
