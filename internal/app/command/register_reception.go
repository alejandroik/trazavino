package command

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/alejandroik/trazavino-api/internal/domain/repository"
)

type RegisterReception struct {
	Weight int
	Sugar  int
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
	pr, err := entity.NewProcess(time.Now(), time.Time{}, "", "", entity.TypeReception.String(), 0)
	if err != nil {
		return err
	}
	insertedProcess, err := h.processRepository.AddProcess(ctx, pr)
	if err != nil {
		return err
	}

	rc, err := entity.NewReception(insertedProcess, nil, cmd.Weight, cmd.Sugar)
	if err != nil {
		return err
	}
	_, err = h.receptionRepository.AddReception(ctx, rc)
	if err != nil {
		return err
	}

	return nil
}
