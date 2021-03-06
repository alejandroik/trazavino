package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type ReceptionRepository struct {
	db *sqlx.DB
}

func NewReceptionRepository(db *sqlx.DB) *ReceptionRepository {
	if db == nil {
		panic("missing db")
	}

	return &ReceptionRepository{db: db}
}

func (r ReceptionRepository) AddReception(ctx context.Context, rc *entity.Reception) (*entity.Reception, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	q := generated.New(tx)

	rm, err := q.AddReception(ctx, generated.AddReceptionParams{
		ID:          rc.UUID(),
		CreatedAt:   time.Now(),
		Weight:      rc.Weight(),
		Sugar:       rc.Sugar(),
		TruckID:     rc.TruckUUID(),
		VineyardID:  rc.VineyardUUID(),
		GrapeTypeID: rc.GrapeTypeUUID(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	insertedReception, err := unmarshalReception(rm)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return insertedReception, tx.Commit()
}

func (r ReceptionRepository) GetReception(ctx context.Context, id int64) (*entity.Reception, error) {
	q := generated.New(r.db)
	rm, err := q.GetReception(ctx, id)
	if err != nil {
		return nil, err
	}

	reception, err := unmarshalReception(rm)
	if err != nil {
		return nil, err
	}

	return reception, nil
}

func (r ReceptionRepository) ListReceptions(context.Context, int32, int32) ([]*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionRepository) UpdateReception(ctx context.Context, receptionId int64, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error {
	return nil
}

func unmarshalReception(rm generated.Reception) (*entity.Reception, error) {
	return entity.UnmarshalReceptionFromDatabase(rm.ID, rm.TruckID, rm.VineyardID, rm.GrapeTypeID, rm.Weight, rm.Sugar)
}
