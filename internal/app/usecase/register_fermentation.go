package usecase

import (
	"context"
	"github.com/pkg/errors"
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
	tankRepository         repository.TankRepository
}

func NewRegisterFermentationHandler(
	fermentationRepository repository.FermentationRepository,
	tankRepository repository.TankRepository,
	log logger.Interface,
) RegisterFermentationHandler {
	if fermentationRepository == nil {
		panic("nil fermentationRepository")
	}
	if tankRepository == nil {
		panic("nil tankRepository")
	}

	return decorator.ApplyDecorators[RegisterFermentation](
		registerFermentationHandler{
			fermentationRepository: fermentationRepository,
			tankRepository:         tankRepository},
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

	tk, err := h.tankRepository.GetTank(ctx, mc.TankUUID())
	if err != nil {
		return err
	}

	if !tk.IsEmpty() {
		return errors.New("tank is not empty")
	}

	if err = h.fermentationRepository.AddFermentation(ctx, mc); err != nil {
		return err
	}

	return nil
}
