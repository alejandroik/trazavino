package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type WineryRepository struct {
	db *pgx.Conn
}

func NewWineryRepository(db *pgx.Conn) *WineryRepository {
	if db == nil {
		panic("missing db")
	}

	return &WineryRepository{db: db}
}

func (r WineryRepository) AddWinery(ctx context.Context, winery *entity.Winery) error {
	wineryUuid, err := uuid.Parse(winery.ID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddWinery(ctx, generated.AddWineryParams{
		ID:        wineryUuid,
		CreatedAt: time.Now(),
		Name:      winery.Name(),
	}); err != nil {
		return err
	}

	return nil
}

func (r WineryRepository) GetWinery(ctx context.Context, wineryId string) (*entity.Winery, error) {
	wineryUuid, err := uuid.Parse(wineryId)
	if err != nil {
		return nil, err
	}

	q := generated.New(r.db)
	wm, err := q.GetWinery(ctx, wineryUuid)
	if err != nil {
		return nil, err
	}

	winery, err := unmarshalWinery(wm)
	if err != nil {
		return nil, err
	}

	return winery, nil
}

func (r WineryRepository) LIstWineries(ctx context.Context, offset int32, limit int32) ([]*entity.Winery, error) {
	q := generated.New(r.db)
	wms, err := q.ListWinerys(ctx, generated.ListWinerysParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	var wineries []*entity.Winery
	for _, wm := range wms {
		winery, err := unmarshalWinery(wm)
		if err != nil {
			return nil, err
		}

		wineries = append(wineries, winery)
	}

	return wineries, nil
}

func (r WineryRepository) UpdateWinery(
	ctx context.Context,
	wineryId string,
	updateFn func(ctx context.Context, winery *entity.Winery) (*entity.Winery, error),
) error {
	return nil
}

func unmarshalWinery(w generated.Winery) (*entity.Winery, error) {
	return nil, nil
}
