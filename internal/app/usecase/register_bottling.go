package usecase

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/repository"
	"github.com/alejandroik/trazavino/pkg/decorator"
	"github.com/alejandroik/trazavino/pkg/logger"
)

type RegisterBottling struct {
	BottlingUUID      string
	BottlingStartTime time.Time

	WineryUUID string

	CaskUUID string

	WineUUID string

	BottleQty int32
}

type RegisterBottlingHandler decorator.Handler[RegisterBottling]

type registerBottlingHandler struct {
	bottlingRepository repository.BottlingRepository
}

func NewRegisterBottlingHandler(
	bottlingRepository repository.BottlingRepository,
	log logger.Interface,
) RegisterBottlingHandler {
	if bottlingRepository == nil {
		panic("nil bottlingRepository")
	}

	return decorator.ApplyDecorators[RegisterBottling](
		registerBottlingHandler{bottlingRepository: bottlingRepository},
		log,
	)
}

func (h registerBottlingHandler) Handle(ctx context.Context, uc RegisterBottling) (err error) {
	mc, err := entity.NewBottling(
		uc.BottlingUUID,
		uc.BottlingStartTime,
		uc.WineryUUID,
		uc.CaskUUID,
		uc.WineUUID,
		uc.BottleQty,
	)
	if err != nil {
		return err
	}

	if err = h.bottlingRepository.AddBottling(ctx, mc); err != nil {
		return err
	}

	return nil
}
