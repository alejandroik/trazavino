package command

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/repository"
)

type RegisterReception struct {
}

type RegisterReceptionHandler Handler[RegisterReception]

type registerReceptionHandler struct {
	repository repository.ReceptionRepository
}

func NewRegisterReceptionHandler(repository repository.ReceptionRepository) RegisterReceptionHandler {
	return registerReceptionHandler{repository: repository}
}

func (h registerReceptionHandler) Handle(ctx context.Context, cmd RegisterReception) error {
	//rc, err := reception.NewReception()
	//if err != nil {
	//	return nil
	//}

	return nil
}
