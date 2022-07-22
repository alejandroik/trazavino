package sqlc

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/entity/enum/process_type"
	"github.com/google/uuid"
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

func (r ReceptionRepository) AddReception(ctx context.Context, rc *entity.Reception) error {
	processUuid, err := uuid.Parse(rc.UUID())
	if err != nil {
		return err
	}
	truckUuid, err := uuid.Parse(rc.TruckUUID())
	if err != nil {
		return err
	}
	vyUuid, err := uuid.Parse(rc.VineyardUUID())
	if err != nil {
		return err
	}
	gtUuid, err := uuid.Parse(rc.GrapeTypeUUID())
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := generated.New(tx)

	ca := time.Now()
	if err = q.AddProcess(ctx, generated.AddProcessParams{
		ID:        processUuid,
		CreatedAt: ca,
		StartTime: rc.StartTime(),
		PType:     process_type.Reception.String(),
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.AddReception(ctx, generated.AddReceptionParams{
		ID:          processUuid,
		CreatedAt:   ca,
		Weight:      rc.Weight(),
		Sugar:       rc.Sugar(),
		TruckID:     truckUuid,
		VineyardID:  vyUuid,
		GrapeTypeID: gtUuid,
	}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r ReceptionRepository) GetReception(ctx context.Context, id string) (*entity.Reception, error) {
	//q := generated.New(r.db)
	//rm, err := q.GetReception(ctx, id)
	//if err != nil {
	//	return nil, err
	//}
	//
	//reception, err := unmarshalReception(rm)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return reception, nil
	return nil, nil
}

func (r ReceptionRepository) ListReceptions(context.Context, int32, int32) ([]*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionRepository) UpdateReception(ctx context.Context, receptionId string, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error {
	return nil
}

func unmarshalReception(rm generated.Reception) (*entity.Reception, error) {
	//return entity.UnmarshalReceptionFromDatabase(rm.ID, rm.TruckID, rm.VineyardID, rm.GrapeTypeID, rm.Weight, rm.Sugar)
	return nil, nil
}
