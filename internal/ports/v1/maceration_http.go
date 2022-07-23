package v1

import (
	"net/http"
	"time"

	"github.com/alejandroik/trazavino/internal/app/usecase"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h HttpServer) GetMacerations(c *gin.Context) {

}

func (h HttpServer) GetMaceration(c *gin.Context, macerationUUID openapi_types.UUID) {

}

func (h HttpServer) RegisterMaceration(c *gin.Context) {
	postMaceration := PostMaceration{}
	if err := c.ShouldBindJSON(&postMaceration); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	uc := usecase.RegisterMaceration{
		MacerationUUID:      uuid.New().String(),
		MacerationStartTime: time.Now().Round(time.Second),
		ReceptionUUID:       postMaceration.ReceptionUuid.String(),
		WarehouseUUID:       postMaceration.WarehouseUuid.String(),
	}

	err := h.app.UseCases.RegisterMaceration.Handle(c, uc)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("content-location", "/macerations/"+uc.MacerationUUID)
}
