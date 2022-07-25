package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
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

func (r GrapeTypeRepository) AddGrapeType(ctx context.Context, grapeType *entity.GrapeType) error {
	gtUuid, err := uuid.Parse(grapeType.ID())
	if err != nil {
		return err
	}
	wineryUuid, err := uuid.Parse(grapeType.WineryUUID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddGrapeType(ctx, generated.AddGrapeTypeParams{
		ID:        gtUuid,
		CreatedAt: time.Now(),
		Name:      grapeType.Name(),
		WineryID:  wineryUuid,
	}); err != nil {
		return err
	}

	return nil
}

func (r GrapeTypeRepository) GetGrapeType(ctx context.Context, grapeTypeId int64) (*entity.GrapeType, error) {
	//q := generated.New(r.db)
	//gtm, err := q.GetGrapeType(ctx, grapeTypeId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//grapeType, err := unmarshalGrapeType(gtm)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return grapeType, nil
	return nil, nil
}

func unmarshalGrapeType(gtm generated.GrapeType) (*entity.GrapeType, error) {
	//return entity.UnmarshalGrapeTypeFromDatabase(gtm.ID, gtm.Name)
	return nil, nil
}
