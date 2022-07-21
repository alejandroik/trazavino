package v1

import "github.com/gin-gonic/gin"

func (h HttpServer) Status(c *gin.Context) {
	c.JSON(200, Status{Message: "ok"})
}
