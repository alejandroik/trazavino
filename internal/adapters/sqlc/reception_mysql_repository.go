package sqlc

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type ReceptionMysqlRepository struct {
	db *sqlx.DB
}

func NewReceptionMysqlRepository(db *sqlx.DB) *ReceptionMysqlRepository {
	if db == nil {
		panic("missing db")
	}

	return &ReceptionMysqlRepository{db: db}
}

func (r ReceptionMysqlRepository) AddReception(ctx context.Context, rc *entity.Reception) (*entity.Reception, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	q := generated.New(tx)

	_, err = q.AddReception(ctx, generated.AddReceptionParams{
		ID:     int64(rc.Process().Id()),
		Weight: int32(rc.Weight()),
		Sugar:  int32(rc.Sugar()),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return rc, tx.Commit()
}

func (r ReceptionMysqlRepository) GetReception(ctx context.Context, receptionId int) (*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionMysqlRepository) GetAllReceptions() ([]*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionMysqlRepository) UpdateReception(ctx context.Context, receptionId int, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error {
	return nil
}
