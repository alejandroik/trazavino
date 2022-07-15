package command

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/alejandroik/trazavino-api/internal/domain/entity/enum/process_type"
	"github.com/alejandroik/trazavino-api/internal/domain/repository"
)

type RegisterMaceration struct {
	ReceptionID int64
	WarehouseID int64
}

type RegisterMacerationHandler Handler[RegisterMaceration]

type registerMacerationHandler struct {
	processRepository    repository.ProcessRepository
	macerationRepository repository.MacerationRepository
}

func NewRegisterMacerationHandler(processRepository repository.ProcessRepository, macerationRepository repository.MacerationRepository) RegisterMacerationHandler {
	if processRepository == nil {
		panic("nil processRepository")
	}
	if macerationRepository == nil {
		panic("nil macerationRepository")
	}

	return registerMacerationHandler{
		processRepository:    processRepository,
		macerationRepository: macerationRepository,
	}
}

func (h registerMacerationHandler) Handle(ctx context.Context, cmd RegisterMaceration) error {
	pr, err := entity.NewProcess(0, time.Now(), time.Time{}, "", "", process_type.Maceration.String(), 0)
	if err != nil {
		return err
	}
	insertedProcess, err := h.processRepository.AddProcess(ctx, pr)
	if err != nil {
		return err
	}

	rc, err := entity.NewMaceration(insertedProcess.ID(), cmd.ReceptionID, cmd.WarehouseID)
	if err != nil {
		return err
	}
	_, err = h.macerationRepository.AddMaceration(ctx, rc)
	if err != nil {
		return err
	}

	return nil
}
