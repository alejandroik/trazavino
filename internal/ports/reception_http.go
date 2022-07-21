package ports

import (
	"net/http"
	"time"

	"github.com/alejandroik/trazavino/internal/app/command"
	"github.com/alejandroik/trazavino/internal/app/query"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h HttpServer) GetReceptions(c *gin.Context) {

}

func (h HttpServer) GetReception(c *gin.Context, receptionUUID openapi_types.UUID) {
	reception, err := h.app.Queries.ReceptionByID.Handle(c, query.ReceptionByID{ReceptionUUID: receptionUUID.String()})
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(200, receptionToResponse(reception))
}

func (h HttpServer) RegisterReception(c *gin.Context) {
	postReception := PostReception{}
	if err := c.ShouldBindJSON(&postReception); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd := command.RegisterReception{
		ReceptionUUID:      uuid.New().String(),
		ReceptionStartTime: time.Now().Round(time.Second),
		TruckUUID:          postReception.TruckUuid.String(),
		TruckLicense:       postReception.Truck,
		VineyardUUID:       postReception.VineyardkUuid.String(),
		VineyardName:       postReception.Vineyard,
		GrapeTypeUUID:      postReception.GrapeTypeUuid.String(),
		GrapeTypeName:      postReception.GrapeType,
		Weight:             postReception.Weight,
		Sugar:              postReception.Sugar,
	}

	err := h.app.Commands.RegisterReception.Handle(c, cmd)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("content-location", "/receptions/"+cmd.ReceptionUUID)
}

func receptionToResponse(rc query.Reception) Reception {
	rUuid, _ := uuid.Parse(rc.UUID)
	tUuid, _ := uuid.Parse(rc.TruckUUID)
	vyUuid, _ := uuid.Parse(rc.VineyardUUID)
	gtUuid, _ := uuid.Parse(rc.GrapeTypeUUID)

	return Reception{
		EndTime:       rc.EndTime,
		GrapeType:     rc.GrapeType,
		GrapeTypeUuid: gtUuid,
		Hash:          rc.Hash,
		StartTime:     rc.StartTime,
		Sugar:         rc.Sugar,
		Transaction:   rc.Transaction,
		Truck:         rc.Truck,
		TruckUuid:     tUuid,
		Uuid:          rUuid,
		Vineyard:      rc.Vineyard,
		VineyardUuid:  vyUuid,
		Weight:        rc.Weight,
	}
}
