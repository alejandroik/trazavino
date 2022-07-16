package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino-api/internal/domain/entity"
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

func (r MacerationRepository) AddMaceration(ctx context.Context, m *entity.Maceration) (*entity.Maceration, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	q := generated.New(tx)

	mm, err := q.AddMaceration(ctx, generated.AddMacerationParams{
		ID:          m.UUID(),
		CreatedAt:   time.Now(),
		ReceptionID: m.ReceptionUUID(),
		WarehouseID: m.WarehouseUUID(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	insertedMaceration, err := unmarshalMaceration(mm)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return insertedMaceration, tx.Commit()
}

func (r MacerationRepository) GetMaceration(ctx context.Context, macerationId int64) (*entity.Maceration, error) {
	q := generated.New(r.db)
	mm, err := q.GetMaceration(ctx, macerationId)
	if err != nil {
		return nil, err
	}

	maceration, err := unmarshalMaceration(mm)
	if err != nil {
		return nil, err
	}

	return maceration, nil
}

func (r MacerationRepository) UpdateMaceration(ctx context.Context, macerationId int64, updateFn func(ctx context.Context, m *entity.Maceration) (*entity.Maceration, error)) {

}

func unmarshalMaceration(mm generated.Maceration) (*entity.Maceration, error) {
	return entity.UnmarshalMacerationFromDatabase(mm.ID, mm.ReceptionID, mm.WarehouseID)
}
