package v1

import (
	"net/http"
	"time"

	"github.com/alejandroik/trazavino/internal/app/usecase"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h HttpServer) GetReceptions(c *gin.Context) {

}

func (h HttpServer) GetReception(c *gin.Context, receptionUUID openapi_types.UUID) {

}

func (h HttpServer) RegisterReception(c *gin.Context) {
	postReception := PostReception{}
	if err := c.ShouldBindJSON(&postReception); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	uc := usecase.RegisterReception{
		ReceptionUUID:      uuid.New().String(),
		ReceptionStartTime: time.Now().Round(time.Second),
		TruckUUID:          postReception.TruckUuid.String(),
		VineyardUUID:       postReception.VineyardkUuid.String(),
		GrapeTypeUUID:      postReception.GrapeTypeUuid.String(),
		Weight:             postReception.Weight,
		Sugar:              postReception.Sugar,
	}

	err := h.app.UseCases.RegisterReception.Handle(c, uc)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("content-location", "/receptions/"+uc.ReceptionUUID)
}
