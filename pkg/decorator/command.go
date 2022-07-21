package decorator

import (
	"context"

	"github.com/alejandroik/trazavino/pkg/logger"
)

func ApplyCommandDecorators[H any](handler CommandHandler[H], log logger.Interface) CommandHandler[H] {
	return commandLoggingDecorator[H]{
		base: handler,
		log:  log,
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}
