package decorator

import (
	"context"

	"github.com/alejandroik/trazavino/pkg/logger"
)

func ApplyDecorators[H any](handler Handler[H], log logger.Interface) Handler[H] {
	return loggingDecorator[H]{
		base: handler,
		log:  log,
	}
}

type Handler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}
