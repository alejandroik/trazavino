package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
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

func (r VineyardRepository) AddVineyard(ctx context.Context, vineyard *entity.Vineyard) error {
	vyUuid, err := uuid.Parse(vineyard.ID())
	if err != nil {
		return err
	}
	wineryUuid, err := uuid.Parse(vineyard.WineryUUID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddVineyard(ctx, generated.AddVineyardParams{
		ID:        vyUuid,
		CreatedAt: time.Now(),
		Name:      vineyard.Name(),
		WineryID:  wineryUuid,
	}); err != nil {
		return err
	}

	return nil
}

func (r VineyardRepository) GetVineyard(ctx context.Context, vineyardId string) (*entity.Vineyard, error) {
	//q := generated.New(r.db)
	//vm, err := q.GetVineyard(ctx, vineyardId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//vineyard, err := unmarshalVineyard(vm)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return vineyard, nil
	return nil, nil
}

func unmarshalVineyard(vm generated.Vineyard) (*entity.Vineyard, error) {
	//return entity.UnmarshalVineyardFromDatabase(vm.ID, vm.Name)
	return nil, nil
}
