package command

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/repository"
)

type RegisterMaceration struct {
	ReceptionID int64
	WarehouseID int64
}

type RegisterMacerationHandler Handler[RegisterMaceration]

type registerMacerationHandler struct {
	macerationRepository repository.MacerationRepository
	receptionRepository  repository.ReceptionRepository
	warehouseRepository  repository.WarehouseRepository
}

func NewRegisterMacerationHandler(receptionRepository repository.ReceptionRepository, macerationRepository repository.MacerationRepository, warehouseRepository repository.WarehouseRepository) RegisterMacerationHandler {
	if macerationRepository == nil {
		panic("nil macerationRepository")
	}
	if receptionRepository == nil {
		panic("nil receptionRepository")
	}
	if warehouseRepository == nil {
		panic("nil warehouseRepository")
	}

	return registerMacerationHandler{
		macerationRepository: macerationRepository,
		receptionRepository:  receptionRepository,
		warehouseRepository:  warehouseRepository,
	}
}

func (h registerMacerationHandler) Handle(ctx context.Context, cmd RegisterMaceration) error {
	return nil
}
