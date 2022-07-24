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
		ReceptionUUID:      uuid.NewString(),
		ReceptionStartTime: time.Now(),
		TruckUUID:          postReception.TruckUuid.String(),
		VineyardUUID:       postReception.VineyardUuid.String(),
		GrapeTypeUUID:      postReception.GrapeTypeUuid.String(),
		Weight:             postReception.Weight,
		Sugar:              postReception.Sugar,
	}

	if err := h.app.UseCases.RegisterReception.Handle(c, uc); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("content-location", "/receptions/"+uc.ReceptionUUID)
}
