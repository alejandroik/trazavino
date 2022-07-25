package v1

import (
	"github.com/alejandroik/trazavino/internal/app/usecase"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (h HttpServer) GetFermentations(c *gin.Context) {

}

func (h HttpServer) RegisterFermentation(c *gin.Context) {
	postFermentation := PostFermentation{}
	if err := c.ShouldBindJSON(&postFermentation); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	uc := usecase.RegisterFermentation{
		FermentationUUID:      uuid.NewString(),
		FermentationStartTime: time.Now(),
		WineryUUID:            postFermentation.WineryUuid.String(),
		WarehouseUUID:         postFermentation.WarehouseUuid.String(),
		TankUUID:              postFermentation.TankUuid.String(),
	}

	if err := h.app.UseCases.RegisterFermentation.Handle(c, uc); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("content-location", "/fermentations/"+uc.FermentationUUID)
}

func (h HttpServer) GetFermentation(c *gin.Context, fermentationUUID openapi_types.UUID) {

}
