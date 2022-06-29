package service

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/app"
)

// TODO implement
func NewApplication(ctx context.Context) (app.Application, func()) {
	return app.Application{}, func() {}
}
