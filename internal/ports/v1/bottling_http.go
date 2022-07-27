package v1

import (
	"net/http"
	"time"

	"github.com/alejandroik/trazavino/internal/app/usecase"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h HttpServer) GetBottlings(c *gin.Context) {

}

func (h HttpServer) RegisterBottling(c *gin.Context) {
	postBottling := PostBottling{}
	if err := c.ShouldBindJSON(&postBottling); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	uc := usecase.RegisterBottling{
		BottlingUUID:      uuid.NewString(),
		BottlingStartTime: time.Now(),
		WineryUUID:        postBottling.WineryUuid.String(),
		CaskUUID:          postBottling.CaskUuid.String(),
		WineUUID:          postBottling.WineUuid.String(),
		BottleQty:         postBottling.BottleQty,
	}

	if err := h.app.UseCases.RegisterBottling.Handle(c, uc); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("content-location", "/bottlings/"+uc.BottlingUUID)
}

func (h HttpServer) GetBottling(c *gin.Context, bottlingUUID openapi_types.UUID) {

}
