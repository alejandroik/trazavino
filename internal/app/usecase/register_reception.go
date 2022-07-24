package usecase

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/repository"
	"github.com/alejandroik/trazavino/pkg/decorator"
	"github.com/alejandroik/trazavino/pkg/logger"
)

type RegisterReception struct {
	ReceptionUUID      string
	ReceptionStartTime time.Time

	WineryUUID string

	TruckUUID string

	VineyardUUID string

	GrapeTypeUUID string

	Weight int32
	Sugar  int32
}

type RegisterReceptionHandler decorator.Handler[RegisterReception]

type registerReceptionHandler struct {
	repository repository.ReceptionRepository
}

func NewRegisterReceptionHandler(
	repository repository.ReceptionRepository,
	log logger.Interface,
) RegisterReceptionHandler {
	if repository == nil {
		panic("nil repository")
	}

	return decorator.ApplyDecorators[RegisterReception](
		registerReceptionHandler{repository: repository},
		log,
	)
}

func (h registerReceptionHandler) Handle(ctx context.Context, cmd RegisterReception) (err error) {
	rc, err := entity.NewReception(
		cmd.ReceptionUUID,
		cmd.ReceptionStartTime,
		cmd.WineryUUID,
		cmd.TruckUUID,
		cmd.VineyardUUID,
		cmd.GrapeTypeUUID,
		cmd.Weight,
		cmd.Sugar)
	if err != nil {
		return err
	}

	if err = h.repository.AddReception(ctx, rc); err != nil {
		return err
	}

	return nil
}
