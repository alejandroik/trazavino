package query

import (
	"context"

	"github.com/alejandroik/trazavino/pkg/decorator"
)

type ReceptionByID struct {
	ReceptionUUID string
}

type ReceptionByIDHandler decorator.QueryHandler[ReceptionByID, Reception]

type receptionByIDHandler struct {
	readModel ReceptionByIDReadModel
}

func NewReceptionByIDHandler(readModel ReceptionByIDReadModel) ReceptionByIDHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return receptionByIDHandler{readModel: readModel}
}

type ReceptionByIDReadModel interface {
	FindReceptionByID(ctx context.Context, receptionUUID string) (Reception, error)
}

func (h receptionByIDHandler) Handle(ctx context.Context, query ReceptionByID) (Reception, error) {
	return h.readModel.FindReceptionByID(ctx, query.ReceptionUUID)
}
