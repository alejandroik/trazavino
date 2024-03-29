package service

import (
	"context"

	"github.com/alejandroik/trazavino/internal/app"
	"github.com/alejandroik/trazavino/internal/app/usecase"
	"github.com/alejandroik/trazavino/pkg/logger"
)

func NewApplication(ctx context.Context, log logger.Logger) app.Application {
	return newApplication(ctx, log)
}

func newApplication(ctx context.Context, log logger.Logger) app.Application {
	r, err := initRepositories(ctx)
	if err != nil {
		panic(err)
	}

	return app.Application{
		UseCases: app.UseCases{
			RegisterReception:    usecase.NewRegisterReceptionHandler(r.ReceptionRepository, log),
			RegisterMaceration:   usecase.NewRegisterMacerationHandler(r.MacerationRepository, r.WarehouseRepository, log),
			RegisterFermentation: usecase.NewRegisterFermentationHandler(r.FermentationRepository, r.TankRepository, log),
			RegisterAgeing:       usecase.NewRegisterAgeingHandler(r.AgeingRepository, r.CaskRepository, log),
			RegisterBottling:     usecase.NewRegisterBottlingHandler(r.BottlingRepository, log),
		},
	}
}
