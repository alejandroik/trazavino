package app

import (
	"github.com/alejandroik/trazavino/internal/app/usecase"
)

type Application struct {
	UseCases UseCases
}

type UseCases struct {
	RegisterReception    usecase.RegisterReceptionHandler
	RegisterMaceration   usecase.RegisterMacerationHandler
	RegisterFermentation usecase.RegisterFermentationHandler
}
