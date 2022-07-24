package usecase

import (
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/repository"
	"github.com/alejandroik/trazavino/pkg/decorator"
	"github.com/alejandroik/trazavino/pkg/logger"
)

type RegisterAgeing struct {
	AgeingUUID string

	AgeingStartTime time.Time

	TankUUID string

	CaskUUID string
}

type RegisterAgeingHandler decorator.Handler[RegisterAgeing]

type registerAgeingHandler struct {
	ageingRepository repository.AgeingRepository
	caskRepository   repository.CaskRepository
}

func NewRegisterAgeingHandler(
	ageingRepository repository.AgeingRepository,
	caskRepository repository.CaskRepository,
	log logger.Interface,
) RegisterAgeingHandler {
	if ageingRepository == nil {
		panic("nil ageingRepository")
	}
	if caskRepository == nil {
		panic("nil caskRepository")
	}

	return decorator.ApplyDecorators[RegisterAgeing](
		registerAgeingHandler{
			ageingRepository: ageingRepository,
			caskRepository:   caskRepository},
		log,
	)
}

func (h registerAgeingHandler) Handle(ctx context.Context, cmd RegisterAgeing) (err error) {
	mc, err := entity.NewAgeing(
		cmd.AgeingUUID,
		cmd.AgeingStartTime,
		cmd.TankUUID,
		cmd.CaskUUID,
	)
	if err != nil {
		return err
	}

	ck, err := h.caskRepository.GetCask(ctx, mc.CaskUUID())
	if err != nil {
		return err
	}

	if !ck.IsEmpty() {
		return errors.New("cask is not empty")
	}

	if err = h.ageingRepository.AddAgeing(ctx, mc); err != nil {
		return err
	}

	return nil
}
