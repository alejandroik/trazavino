package main

import (
	"context"
	"net/http"

	"github.com/alejandroik/trazavino-api/internal/ports"
	"github.com/alejandroik/trazavino-api/internal/service"
	"github.com/alejandroik/trazavino-api/pkg/server"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	app, shutdown := service.NewApplication(ctx)
	defer shutdown()

	ports.NewHttpServer(app)

	//TODO implement
	server.RunHTTPServer(func(router *gin.Engine) http.Handler {
		return gin.Default()
	})
}
