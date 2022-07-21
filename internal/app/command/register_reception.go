package command

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/repository"
	"github.com/alejandroik/trazavino/pkg/decorator"
	"github.com/alejandroik/trazavino/pkg/logger"
)

type RegisterReception struct {
	ReceptionUUID string

	ReceptionStartTime time.Time

	TruckUUID    string
	TruckLicense string

	VineyardUUID string
	VineyardName string

	GrapeTypeUUID string
	GrapeTypeName string

	Weight int32
	Sugar  int32
}

type RegisterReceptionHandler decorator.CommandHandler[RegisterReception]

type registerReceptionHandler struct {
	receptionRepository repository.ReceptionRepository
}

func NewRegisterReceptionHandler(
	receptionRepository repository.ReceptionRepository,
	log logger.Interface,
) RegisterReceptionHandler {
	if receptionRepository == nil {
		panic("nil receptionRepository")
	}

	return decorator.ApplyCommandDecorators[RegisterReception](
		registerReceptionHandler{receptionRepository: receptionRepository},
		log,
	)
}

func (h registerReceptionHandler) Handle(ctx context.Context, cmd RegisterReception) (err error) {
	rc, err := entity.NewReception(
		cmd.ReceptionUUID,
		cmd.ReceptionStartTime,
		cmd.TruckUUID,
		cmd.TruckLicense,
		cmd.VineyardUUID,
		cmd.VineyardName,
		cmd.GrapeTypeUUID,
		cmd.GrapeTypeName,
		cmd.Weight,
		cmd.Sugar)
	if err != nil {
		return err
	}

	if err = h.receptionRepository.AddReception(ctx, rc); err != nil {
		return err
	}

	return nil
}
