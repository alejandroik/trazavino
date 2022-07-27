package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type WineRepository struct {
	db *pgx.Conn
}

func NewWineRepository(db *pgx.Conn) *WineRepository {
	if db == nil {
		panic("missing db")
	}

	return &WineRepository{db: db}
}

func (r WineRepository) AddWine(ctx context.Context, wine *entity.Wine) error {
	wUuid, err := uuid.Parse(wine.ID())
	if err != nil {
		return err
	}
	wineryUuid, err := uuid.Parse(wine.WineryUUID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddWine(ctx, generated.AddWineParams{
		ID:        wUuid,
		CreatedAt: time.Now(),
		Name:      wine.Name(),
		WineryID:  wineryUuid,
	}); err != nil {
		return err
	}

	return nil
}

func (r WineRepository) GetWine(ctx context.Context, wineId string) (*entity.Wine, error) {
	wUuid, err := uuid.Parse(wineId)
	if err != nil {
		return nil, err
	}

	q := generated.New(r.db)
	cm, err := q.GetWine(ctx, wUuid)
	if err != nil {
		return nil, err
	}

	wine, err := unmarshalWine(cm)
	if err != nil {
		return nil, err
	}

	return wine, nil
}

func (r WineRepository) ListWines(ctx context.Context, offset int32, limit int32) ([]*entity.Wine, error) {
	q := generated.New(r.db)
	wms, err := q.ListWines(ctx, generated.ListWinesParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	var wines []*entity.Wine
	for _, wm := range wms {
		wine, err := unmarshalWine(wm)
		if err != nil {
			return nil, err
		}

		wines = append(wines, wine)
	}

	return wines, nil
}

func (r WineRepository) UpdateWine(
	ctx context.Context,
	wineId string,
	updateFn func(ctx context.Context, wine *entity.Wine) (*entity.Wine, error),
) error {
	return nil
}

func unmarshalWine(w generated.Wine) (*entity.Wine, error) {
	return nil, nil
}
