package command

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/entity/enum/process_type"
	"github.com/alejandroik/trazavino/internal/domain/repository"
)

type RegisterMaceration struct {
	ReceptionID int64
	WarehouseID int64
}

type RegisterMacerationHandler Handler[RegisterMaceration]

type registerMacerationHandler struct {
	processRepository    repository.ProcessRepository
	macerationRepository repository.MacerationRepository
	warehouseRepository  repository.WarehouseRepository
}

func NewRegisterMacerationHandler(processRepository repository.ProcessRepository, macerationRepository repository.MacerationRepository, warehouseRepository repository.WarehouseRepository) RegisterMacerationHandler {
	if processRepository == nil {
		panic("nil processRepository")
	}
	if macerationRepository == nil {
		panic("nil macerationRepository")
	}
	if warehouseRepository == nil {
		panic("nil warehouseRepository")
	}

	return registerMacerationHandler{
		processRepository:    processRepository,
		macerationRepository: macerationRepository,
		warehouseRepository:  warehouseRepository,
	}
}

func (h registerMacerationHandler) Handle(ctx context.Context, cmd RegisterMaceration) error {
	date := time.Now()

	pr, err := entity.NewProcess(0, date, time.Time{}, "", "", process_type.Maceration.String(), 0)
	if err != nil {
		return err
	}
	insertedProcess, err := h.processRepository.AddProcess(ctx, pr)
	if err != nil {
		return err
	}

	rc, err := entity.NewMaceration(insertedProcess.UUID(), cmd.ReceptionID, cmd.WarehouseID)
	if err != nil {
		return err
	}
	mac, err := h.macerationRepository.AddMaceration(ctx, rc)
	if err != nil {
		return err
	}

	// set warehouse as occupied
	if err = h.warehouseRepository.UpdateWarehouse(
		ctx,
		cmd.WarehouseID,
		func(ctx context.Context, wh *entity.Warehouse) (*entity.Warehouse, error) {
			_ = wh.UpdateIsEmpty(false)

			return wh, nil
		}); err != nil {
		return err
	}

	// update reception end date
	if err = h.processRepository.UpdateProcess(
		ctx,
		cmd.ReceptionID,
		func(ctx context.Context, pr *entity.Process) (*entity.Process, error) {
			_ = pr.UpdateEndDate(date)

			return pr, nil
		}); err != nil {
		return err
	}

	// set reception as previous process
	if err = h.processRepository.UpdateProcess(
		ctx,
		mac.UUID(),
		func(ctx context.Context, pr *entity.Process) (*entity.Process, error) {
			_ = pr.UpdatePreviousUUID(cmd.ReceptionID)

			return pr, nil
		}); err != nil {
		return err
	}

	return nil
}
