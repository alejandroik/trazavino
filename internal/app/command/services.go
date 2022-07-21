package command

import "context"

type ReceptionService interface {
	RegisterReception(ctx context.Context) error
}

type MacerationService interface {
	RegisterMaceration(ctx context.Context) error
}
