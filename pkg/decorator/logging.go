package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/alejandroik/trazavino/pkg/logger"
)

type commandLoggingDecorator[C any] struct {
	base CommandHandler[C]
	log  logger.Interface
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	defer func() {
		if err == nil {
			d.log.Infow("Command executed successfully", fields(cmd)...)
		} else {
			d.log.Errorw("Failed to execute command", fieldsWithError(cmd, err)...)
		}
	}()

	return d.base.Handle(ctx, cmd)
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}

func fields(cmd interface{}) []interface{} {
	return []interface{}{
		"command", generateActionName(cmd),
		"command_body", fmt.Sprintf("%+v", cmd),
	}
}

func fieldsWithError(cmd interface{}, err error) []interface{} {
	return append(fields(cmd), "error", err)
}
