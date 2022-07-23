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

type FermentationRepository struct {
	db *sqlx.DB
}

func NewFermentationRepository(db *sqlx.DB) *FermentationRepository {
	if db == nil {
		panic("missing db")
	}

	return &FermentationRepository{db: db}
}

func (r FermentationRepository) AddFermentation(ctx context.Context, f *entity.Fermentation) error {
	feUuid, err := uuid.Parse(f.UUID())
	if err != nil {
		return nil
	}
	whUuid, err := uuid.Parse(f.WarehouseUUID())
	if err != nil {
		return err
	}
	tankUuid, err := uuid.Parse(f.TankUUID())
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := generated.New(tx)

	ca := time.Now()
	if err = q.AddProcess(ctx, generated.AddProcessParams{
		ID:        feUuid,
		CreatedAt: ca,
		StartTime: f.StartTime(),
		PType:     process_type.Maceration.String(),
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.AddFermentation(ctx, generated.AddFermentationParams{
		ID:          feUuid,
		CreatedAt:   ca,
		WarehouseID: whUuid,
		TankID:      tankUuid,
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.UpdateProcess(ctx, generated.UpdateProcessParams{
		//ID:        ,
		UpdatedAt: sql.NullTime{Time: ca, Valid: true},
		EndTime:   sql.NullTime{Time: f.StartTime(), Valid: true},
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.UpdateWarehouse(ctx, generated.UpdateWarehouseParams{
		ID:        whUuid,
		UpdatedAt: sql.NullTime{Time: ca, Valid: true},
		IsEmpty:   false,
	}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r FermentationRepository) GetFermentation(ctx context.Context, fermentationUUID string) (*entity.Fermentation, error) {
	return nil, nil
}

func (r FermentationRepository) UpdateFermentation(
	ctx context.Context,
	fermentationUUID string,
	updateFn func(ctx context.Context, f *entity.Fermentation) (*entity.Fermentation, error),
) error {
	return nil
}
