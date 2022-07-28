package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/entity/enum/process_type"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type FermentationRepository struct {
	db *pgx.Conn
}

func NewFermentationRepository(db *pgx.Conn) *FermentationRepository {
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
	wineryUuid, err := uuid.Parse(f.WineryUUID())
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

	mac, err := q.FindMaceration(ctx, whUuid)
	if err != nil {
		return err
	}

	if err = q.UpdateProcess(ctx, generated.UpdateProcessParams{
		ID:        mac.ID,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		EndTime:   sql.NullTime{Time: f.StartTime(), Valid: true},
	}); err != nil {
		return err
	}

	if err = q.UpdateWarehouseUsage(ctx, generated.UpdateWarehouseUsageParams{
		ID:        whUuid,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		IsEmpty:   true,
	}); err != nil {
		return err
	}

	if err = q.UpdateTankUsage(ctx, generated.UpdateTankUsageParams{
		ID:        tankUuid,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		IsEmpty:   false,
	}); err != nil {
		return err
	}

	if err = q.AddProcess(ctx, generated.AddProcessParams{
		ID:         feUuid,
		CreatedAt:  now,
		StartTime:  f.StartTime(),
		PType:      process_type.Fermentation.String(),
		PreviousID: uuid.NullUUID{UUID: mac.ID, Valid: true},
		WineryID:   wineryUuid,
	}); err != nil {
		return err
	}

	if err = q.AddFermentation(ctx, generated.AddFermentationParams{
		ID:          feUuid,
		CreatedAt:   now,
		WarehouseID: whUuid,
		TankID:      tankUuid,
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
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
