package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/entity/enum/process_type"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MacerationRepository struct {
	db *sqlx.DB
}

func NewMacerationRepository(db *sqlx.DB) *MacerationRepository {
	if db == nil {
		panic("missing db")
	}

	return &MacerationRepository{db: db}
}

func (r MacerationRepository) AddMaceration(ctx context.Context, m *entity.Maceration) error {
	mcUuid, err := uuid.Parse(m.UUID())
	if err != nil {
		return err
	}
	recUuid, err := uuid.Parse(m.ReceptionUUID())
	if err != nil {
		return err
	}
	whUuid, err := uuid.Parse(m.WarehouseUUID())
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := generated.New(tx)

	now := time.Now()

	if err = q.UpdateProcess(ctx, generated.UpdateProcessParams{
		ID:        recUuid,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		EndTime:   sql.NullTime{Time: m.StartTime(), Valid: true},
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.UpdateWarehouse(ctx, generated.UpdateWarehouseParams{
		ID:        whUuid,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		IsEmpty:   false,
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.AddProcess(ctx, generated.AddProcessParams{
		ID:         mcUuid,
		CreatedAt:  now,
		StartTime:  m.StartTime(),
		PType:      process_type.Maceration.String(),
		PreviousID: uuid.NullUUID{UUID: recUuid, Valid: true},
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.AddMaceration(ctx, generated.AddMacerationParams{
		ID:          mcUuid,
		CreatedAt:   now,
		ReceptionID: recUuid,
		WarehouseID: whUuid,
	}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r MacerationRepository) GetMaceration(ctx context.Context, macerationId string) (*entity.Maceration, error) {
	return nil, nil
}

func (r MacerationRepository) UpdateMaceration(
	ctx context.Context,
	macerationId string,
	updateFn func(ctx context.Context, m *entity.Maceration) (*entity.Maceration, error),
) error {
	return nil
}

func unmarshalMaceration(mm generated.Maceration) (*entity.Maceration, error) {
	return nil, nil
}
