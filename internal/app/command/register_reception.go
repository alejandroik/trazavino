package command

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/repository"
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

type RegisterReceptionHandler Handler[RegisterReception]

type registerReceptionHandler struct {
	receptionRepository repository.ReceptionRepository
}

func NewRegisterReceptionHandler(receptionRepository repository.ReceptionRepository) RegisterReceptionHandler {
	if receptionRepository == nil {
		panic("nil receptionRepository")
	}

	return registerReceptionHandler{
		receptionRepository: receptionRepository,
	}
}

func (h registerReceptionHandler) Handle(ctx context.Context, cmd RegisterReception) error {
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

	err = h.receptionRepository.AddReception(ctx, rc)
	if err != nil {
		return err
	}

	return nil
}
