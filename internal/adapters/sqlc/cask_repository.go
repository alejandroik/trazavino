package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CaskRepository struct {
	db *sqlx.DB
}

func NewCaskRepository(db *sqlx.DB) *CaskRepository {
	if db == nil {
		panic("missing db")
	}

	return &CaskRepository{db: db}
}

func (r CaskRepository) AddCask(ctx context.Context, cask *entity.Cask) error {
	tkUuid, err := uuid.Parse(cask.ID())
	if err != nil {
		return err
	}
	wineryUuid, err := uuid.Parse(cask.WineryUUID())
	if err != nil {
		return err
	}

	q := generated.New(r.db)

	if err = q.AddCask(ctx, generated.AddCaskParams{
		ID:        tkUuid,
		CreatedAt: time.Now(),
		Name:      cask.Name(),
		CType:     cask.CType(),
		IsEmpty:   cask.IsEmpty(),
		WineryID:  wineryUuid,
	}); err != nil {
		return err
	}

	return nil
}

func (r CaskRepository) GetCask(ctx context.Context, caskId string) (*entity.Cask, error) {
	cUuid, err := uuid.Parse(caskId)
	if err != nil {
		return nil, err
	}

	q := generated.New(r.db)
	cm, err := q.GetCask(ctx, cUuid)
	if err != nil {
		return nil, err
	}

	cask, err := unmarshalCask(cm)
	if err != nil {
		return nil, err
	}

	return cask, nil
}

func (r CaskRepository) ListCasks(ctx context.Context, offset int32, limit int32) ([]*entity.Cask, error) {
	q := generated.New(r.db)
	wms, err := q.ListCasks(ctx, generated.ListCasksParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	var casks []*entity.Cask
	for _, wm := range wms {
		cask, err := unmarshalCask(wm)
		if err != nil {
			return nil, err
		}

		casks = append(casks, cask)
	}

	return casks, nil
}

func (r CaskRepository) UpdateCask(
	ctx context.Context,
	caskId string,
	updateFn func(ctx context.Context, cask *entity.Cask) (*entity.Cask, error),
) error {
	return nil
}

func unmarshalCask(c generated.Cask) (*entity.Cask, error) {
	return nil, nil
}
