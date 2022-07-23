package usecase

import (
	"context"
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
	receptionRepository  repository.ReceptionRepository
	warehouseRepository  repository.WarehouseRepository
}

func NewRegisterMacerationHandler(
	macerationRepository repository.MacerationRepository,
	log logger.Interface,
) RegisterMacerationHandler {
	if macerationRepository == nil {
		panic("nil macerationRepository")
	}

	return decorator.ApplyDecorators[RegisterMaceration](
		registerMacerationHandler{macerationRepository: macerationRepository},
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

	if err = h.macerationRepository.AddMaceration(ctx, mc); err != nil {
		return err
	}

	return nil
}
