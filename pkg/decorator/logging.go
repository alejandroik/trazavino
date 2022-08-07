package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/alejandroik/trazavino/pkg/logger"
)

type loggingDecorator[C any] struct {
	base Handler[C]
	log  logger.Logger
}

func (d loggingDecorator[C]) Handle(ctx context.Context, uc C) (err error) {
	defer func() {
		if err == nil {
			d.log.Infow("Use case executed successfully", fields(uc)...)
		} else {
			d.log.Errorw("Failed to execute use case", fieldsWithError(uc, err)...)
		}
	}()

	return d.base.Handle(ctx, uc)
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}

func fields(uc interface{}) []interface{} {
	return []interface{}{
		"usecase", generateActionName(uc),
		"usecase_body", fmt.Sprintf("%+v", uc),
	}
}

func fieldsWithError(uc interface{}, err error) []interface{} {
	return append(fields(uc), "error", err)
}
