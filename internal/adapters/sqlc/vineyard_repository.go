package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type VineyardRepository struct {
	db *sqlx.DB
}

func NewVineyardRepository(db *sqlx.DB) *VineyardRepository {
	if db == nil {
		panic("missing db")
	}

	return &VineyardRepository{db: db}
}

func (r VineyardRepository) AddVineyard(ctx context.Context, vineyard *entity.Vineyard) (*entity.Vineyard, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	q := generated.New(tx)

	vm, err := q.AddVineyard(ctx, generated.AddVineyardParams{
		CreatedAt: time.Now(),
		Name:      vineyard.Name(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	insertedVineyard, err := unmarshalVineyard(vm)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return insertedVineyard, tx.Commit()
}

func (r VineyardRepository) GetVineyard(ctx context.Context, vineyardId int64) (*entity.Vineyard, error) {
	q := generated.New(r.db)
	vm, err := q.GetVineyard(ctx, vineyardId)
	if err != nil {
		return nil, err
	}

	vineyard, err := unmarshalVineyard(vm)
	if err != nil {
		return nil, err
	}

	return vineyard, nil
}

func unmarshalVineyard(vm generated.Vineyard) (*entity.Vineyard, error) {
	return entity.UnmarshalVineyardFromDatabase(vm.ID, vm.Name)
}
