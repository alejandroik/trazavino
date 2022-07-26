package sqlc

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
)

type ProcessRepository struct {
	db *pgx.Conn
}

func NewProcessRepository(db *pgx.Conn) *ProcessRepository {
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
		pm.WineryID.String(),
		pm.StartTime,
		pm.PType,
		endTime,
		previousUUID,
		hash,
		transaction,
	)
}
