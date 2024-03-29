package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/entity/enum/process_type"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type ReceptionRepository struct {
	db *pgx.Conn
}

func NewReceptionRepository(db *pgx.Conn) *ReceptionRepository {
	if db == nil {
		panic("missing db")
	}

	return &ReceptionRepository{db: db}
}

func (r ReceptionRepository) AddReception(ctx context.Context, rc *entity.Reception) error {
	processUuid, err := uuid.Parse(rc.UUID())
	if err != nil {
		return err
	}
	truckUuid, err := uuid.Parse(rc.TruckUUID())
	if err != nil {
		return err
	}
	vyUuid, err := uuid.Parse(rc.VineyardUUID())
	if err != nil {
		return err
	}
	gtUuid, err := uuid.Parse(rc.GrapeTypeUUID())
	if err != nil {
		return err
	}
	wineryUuid, err := uuid.Parse(rc.WineryUUID())
	if err != nil {
		return err
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	q := generated.New(tx)

	now := time.Now()

	if err = q.AddProcess(ctx, generated.AddProcessParams{
		ID:        processUuid,
		CreatedAt: now,
		StartTime: rc.StartTime(),
		PType:     process_type.Reception.String(),
		WineryID:  wineryUuid,
	}); err != nil {
		return err
	}

	if err = q.AddReception(ctx, generated.AddReceptionParams{
		ID:          processUuid,
		CreatedAt:   now,
		Weight:      rc.Weight(),
		Sugar:       rc.Sugar(),
		TruckID:     truckUuid,
		VineyardID:  vyUuid,
		GrapeTypeID: gtUuid,
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r ReceptionRepository) GetReception(ctx context.Context, id string) (*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionRepository) ListReceptions(context.Context, int32, int32) ([]*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionRepository) UpdateReception(ctx context.Context, receptionId string, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error {
	return nil
}

func unmarshalReception(rm generated.Reception) (*entity.Reception, error) {
	return nil, nil
}
