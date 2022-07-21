package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type GrapeTypeRepository struct {
	db *sqlx.DB
}

func NewGrapeTypeRepository(db *sqlx.DB) *GrapeTypeRepository {
	if db == nil {
		panic("missing db")
	}

	return &GrapeTypeRepository{db: db}
}

func (r GrapeTypeRepository) AddGrapeType(ctx context.Context, grapeType *entity.GrapeType) (*entity.GrapeType, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	q := generated.New(tx)

	gtm, err := q.AddGrapeType(ctx, generated.AddGrapeTypeParams{
		CreatedAt: time.Now(),
		Name:      grapeType.Name(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	insertedGrapeType, err := unmarshalGrapeType(gtm)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return insertedGrapeType, tx.Commit()
}

func (r GrapeTypeRepository) GetGrapeType(ctx context.Context, grapeTypeId int64) (*entity.GrapeType, error) {
	q := generated.New(r.db)
	gtm, err := q.GetGrapeType(ctx, grapeTypeId)
	if err != nil {
		return nil, err
	}

	grapeType, err := unmarshalGrapeType(gtm)
	if err != nil {
		return nil, err
	}

	return grapeType, nil
}

func unmarshalGrapeType(gtm generated.GrapeType) (*entity.GrapeType, error) {
	return entity.UnmarshalGrapeTypeFromDatabase(gtm.ID, gtm.Name)
}
