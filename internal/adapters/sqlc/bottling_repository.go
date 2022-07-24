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

type BottlingRepository struct {
	db *sqlx.DB
}

func NewBottlingRepository(db *sqlx.DB) *BottlingRepository {
	if db == nil {
		panic("missing db")
	}

	return &BottlingRepository{db: db}
}

func (r BottlingRepository) AddBottling(ctx context.Context, a *entity.Bottling) error {
	botUuid, err := uuid.Parse(a.UUID())
	if err != nil {
		return nil
	}
	caskUuid, err := uuid.Parse(a.CaskUUID())
	if err != nil {
		return err
	}
	wineUuid, err := uuid.Parse(a.WineUUID())
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := generated.New(tx)

	now := time.Now()

	age, err := q.FindAgeing(ctx, caskUuid)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = q.UpdateProcess(ctx, generated.UpdateProcessParams{
		ID:        age.ID,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		EndTime:   sql.NullTime{Time: a.StartTime(), Valid: true},
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.UpdateCask(ctx, generated.UpdateCaskParams{
		ID:        caskUuid,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		IsEmpty:   true,
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.AddProcess(ctx, generated.AddProcessParams{
		ID:         botUuid,
		CreatedAt:  now,
		StartTime:  a.StartTime(),
		PType:      process_type.Bottling.String(),
		PreviousID: uuid.NullUUID{UUID: age.ID, Valid: true},
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.AddBottling(ctx, generated.AddBottlingParams{
		ID:        botUuid,
		CreatedAt: now,
		CaskID:    caskUuid,
		WineID:    wineUuid,
		BottleQty: a.BottleQty(),
	}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r BottlingRepository) GetBottling(ctx context.Context, bottlingUUID string) (*entity.Bottling, error) {
	return nil, nil
}

func (r BottlingRepository) UpdateBottling(
	ctx context.Context,
	bottlingUUID string,
	updateFn func(ctx context.Context, b *entity.Bottling) (*entity.Bottling, error),
) error {
	return nil
}
