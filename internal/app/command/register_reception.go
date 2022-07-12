package command

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/repository"
)

type RegisterReception struct {
}

type RegisterReceptionHandler Handler[RegisterReception]

type registerReceptionHandler struct {
	processRepository   repository.ProcessRepository
	receptionRepository repository.ReceptionRepository
}

func NewRegisterReceptionHandler(processRepository repository.ProcessRepository, receptionRepository repository.ReceptionRepository) RegisterReceptionHandler {
	return registerReceptionHandler{
		processRepository:   processRepository,
		receptionRepository: receptionRepository,
	}
}

func (h registerReceptionHandler) Handle(ctx context.Context, cmd RegisterReception) error {
	//rc, err := reception.NewReception()
	//if err != nil {
	//	return nil
	//}

	return nil
}
