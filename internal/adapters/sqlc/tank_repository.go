package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TankRepository struct {
	db *sqlx.DB
}

func NewTankRepository(db *sqlx.DB) *TankRepository {
	if db == nil {
		panic("missing db")
	}

	return &TankRepository{db: db}
}

func (r TankRepository) AddTank(ctx context.Context, tank *entity.Tank) error {
	tkUuid, err := uuid.Parse(tank.ID())
	if err != nil {
		return err
	}
	wineryUuid, err := uuid.Parse(tank.WineryUUID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddTank(ctx, generated.AddTankParams{
		ID:        tkUuid,
		CreatedAt: time.Now(),
		Name:      tank.Name(),
		IsEmpty:   tank.IsEmpty(),
		WineryID:  wineryUuid,
	}); err != nil {
		return err
	}

	return nil
}

func (r TankRepository) GetTank(ctx context.Context, tankId string) (*entity.Tank, error) {
	tkUuid, err := uuid.Parse(tankId)
	if err != nil {
		return nil, err
	}

	q := generated.New(r.db)
	tm, err := q.GetTank(ctx, tkUuid)
	if err != nil {
		return nil, err
	}

	tank, err := unmarshalTank(tm)
	if err != nil {
		return nil, err
	}

	return tank, nil
}

func (r TankRepository) ListTanks(ctx context.Context, offset int32, limit int32) ([]*entity.Tank, error) {
	q := generated.New(r.db)
	wms, err := q.ListTanks(ctx, generated.ListTanksParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	var tanks []*entity.Tank
	for _, wm := range wms {
		tank, err := unmarshalTank(wm)
		if err != nil {
			return nil, err
		}

		tanks = append(tanks, tank)
	}

	return tanks, nil
}

func (r TankRepository) UpdateTank(
	ctx context.Context,
	tankId string,
	updateFn func(ctx context.Context, tank *entity.Tank) (*entity.Tank, error),
) error {
	return nil
}

func unmarshalTank(t generated.Tank) (*entity.Tank, error) {
	return nil, nil
}
