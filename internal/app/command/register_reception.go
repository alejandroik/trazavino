package command

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/alejandroik/trazavino-api/internal/domain/entity/enum/process_type"
	"github.com/alejandroik/trazavino-api/internal/domain/repository"
)

type RegisterReception struct {
	Weight      int32
	Sugar       int32
	TruckID     int64
	VineyardID  int64
	GrapeTypeID int64
}

type RegisterReceptionHandler Handler[RegisterReception]

type registerReceptionHandler struct {
	processRepository   repository.ProcessRepository
	receptionRepository repository.ReceptionRepository
}

func NewRegisterReceptionHandler(processRepository repository.ProcessRepository, receptionRepository repository.ReceptionRepository) RegisterReceptionHandler {
	if processRepository == nil {
		panic("nil processRepository")
	}
	if receptionRepository == nil {
		panic("nil receptionRepository")
	}

	return registerReceptionHandler{
		processRepository:   processRepository,
		receptionRepository: receptionRepository,
	}
}

func (h registerReceptionHandler) Handle(ctx context.Context, cmd RegisterReception) error {
	pr, err := entity.NewProcess(0, time.Now(), time.Time{}, "", "", process_type.Reception.String(), 0)
	if err != nil {
		return err
	}
	insertedProcess, err := h.processRepository.AddProcess(ctx, pr)
	if err != nil {
		return err
	}

	rc, err := entity.NewReception(insertedProcess.ID(), cmd.TruckID, cmd.VineyardID, cmd.GrapeTypeID, cmd.Weight, cmd.Sugar)
	if err != nil {
		return err
	}
	_, err = h.receptionRepository.AddReception(ctx, rc)
	if err != nil {
		return err
	}

	return nil
}
