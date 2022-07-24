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

type RegisterMaceration struct {
	MacerationUUID string

	MacerationStartTime time.Time

	ReceptionUUID string

	WarehouseUUID string
}

type RegisterMacerationHandler decorator.Handler[RegisterMaceration]

type registerMacerationHandler struct {
	macerationRepository repository.MacerationRepository
	warehouseRepository  repository.WarehouseRepository
}

func NewRegisterMacerationHandler(
	macerationRepository repository.MacerationRepository,
	warehouseRepository repository.WarehouseRepository,
	log logger.Interface,
) RegisterMacerationHandler {
	if macerationRepository == nil {
		panic("nil macerationRepository")
	}
	if warehouseRepository == nil {
		panic("nil warehouseRepository")
	}

	return decorator.ApplyDecorators[RegisterMaceration](
		registerMacerationHandler{
			macerationRepository: macerationRepository,
			warehouseRepository:  warehouseRepository},
		log,
	)
}

func (h registerMacerationHandler) Handle(ctx context.Context, cmd RegisterMaceration) (err error) {
	mc, err := entity.NewMaceration(
		cmd.MacerationUUID,
		cmd.MacerationStartTime,
		cmd.ReceptionUUID,
		cmd.WarehouseUUID,
	)
	if err != nil {
		return err
	}

	wh, err := h.warehouseRepository.GetWarehouse(ctx, mc.WarehouseUUID())
	if err != nil {
		return err
	}

	if !wh.IsEmpty() {
		return errors.New("warehouse is not empty")
	}

	if err = h.macerationRepository.AddMaceration(ctx, mc); err != nil {
		return err
	}

	return nil
}
