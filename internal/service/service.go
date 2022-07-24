package service

import (
	"context"

	"github.com/alejandroik/trazavino/internal/app"
	"github.com/alejandroik/trazavino/internal/app/usecase"
	"github.com/alejandroik/trazavino/pkg/logger"
)

func NewApplication(ctx context.Context, log logger.Interface) app.Application {
	return newApplication(ctx, log)
}

func newApplication(ctx context.Context, log logger.Interface) app.Application {
	r, err := getRepositories(ctx)
	if err != nil {
		panic(err)
	}

	return app.Application{
		UseCases: app.UseCases{
			RegisterReception:    usecase.NewRegisterReceptionHandler(r.ReceptionRepository, log),
			RegisterMaceration:   usecase.NewRegisterMacerationHandler(r.MacerationRepository, log),
			RegisterFermentation: usecase.NewRegisterFermentationHandler(r.FermentationRepository, log),
		},
	}
}
