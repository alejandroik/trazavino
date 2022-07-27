package v1

import (
	"net/http"
	"time"

	"github.com/alejandroik/trazavino/internal/app/usecase"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h HttpServer) GetAgeings(c *gin.Context) {

}

func (h HttpServer) RegisterAgeing(c *gin.Context) {
	postAgeing := PostAgeing{}
	if err := c.ShouldBindJSON(&postAgeing); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	uc := usecase.RegisterAgeing{
		AgeingUUID:      uuid.NewString(),
		AgeingStartTime: time.Now(),
		WineryUUID:      postAgeing.WineryUuid.String(),
		TankUUID:        postAgeing.TankUuid.String(),
		CaskUUID:        postAgeing.CaskUuid.String(),
	}

	if err := h.app.UseCases.RegisterAgeing.Handle(c, uc); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("content-location", "/ageings/"+uc.AgeingUUID)
}

func (h HttpServer) GetAgeing(c *gin.Context, ageingUUID openapi_types.UUID) {

}
