package query

import "context"

type Handler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
