package query

import "context"

type ReceptionByID struct {
	ReceptionUUID string
}

type ReceptionByIDHandler Handler[ReceptionByID, Reception]

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