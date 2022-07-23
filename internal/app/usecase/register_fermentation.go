package usecase

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/repository"
	"github.com/alejandroik/trazavino/pkg/decorator"
	"github.com/alejandroik/trazavino/pkg/logger"
)

type RegisterFermentation struct {
	FermentationUUID string

	FermentationStartTime time.Time

	WarehouseUUID string

	TankUUID string
}

type RegisterFermentationHandler decorator.Handler[RegisterFermentation]

type registerFermentationHandler struct {
	fermentationRepository repository.FermentationRepository
}

func NewRegisterFermentationHandler(
	fermentationRepository repository.FermentationRepository,
	log logger.Interface,
) RegisterFermentationHandler {
	if fermentationRepository == nil {
		panic("nil fermentationRepository")
	}

	return decorator.ApplyDecorators[RegisterFermentation](
		registerFermentationHandler{fermentationRepository: fermentationRepository},
		log,
	)
}

func (h registerFermentationHandler) Handle(ctx context.Context, cmd RegisterFermentation) (err error) {
	mc, err := entity.NewFermentation(
		cmd.FermentationUUID,
		cmd.FermentationStartTime,
		cmd.WarehouseUUID,
		cmd.TankUUID,
	)
	if err != nil {
		return err
	}

	if err = h.fermentationRepository.AddFermentation(ctx, mc); err != nil {
		return err
	}

	return nil
}
