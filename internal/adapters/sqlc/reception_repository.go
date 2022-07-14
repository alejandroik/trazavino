package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino-api/internal/domain/entity"
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

	_, err = q.AddReception(ctx, generated.AddReceptionParams{
		ID:        int64(rc.Process().Id()),
		CreatedAt: time.Now(),
		Weight:    int32(rc.Weight()),
		Sugar:     int32(rc.Sugar()),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return rc, tx.Commit()
}

func (r ReceptionRepository) GetReception(ctx context.Context, receptionId int) (*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionRepository) GetAllReceptions() ([]*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionRepository) UpdateReception(ctx context.Context, receptionId int, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error {
	return nil
}
