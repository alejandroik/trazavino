package command

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/reception"
)

type RegisterReception struct {
}

type RegisterReceptionHandler Handler[RegisterReception]

type registerReceptionHandler struct {
	repository reception.Repository
}

func NewRegisterReceptionHandler(repository reception.Repository) RegisterReceptionHandler {
	return registerReceptionHandler{repository: repository}
}

func (h registerReceptionHandler) Handle(ctx context.Context, cmd RegisterReception) error {

	return nil
}
