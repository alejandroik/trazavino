package main

import (
	"context"

	"github.com/alejandroik/trazavino/internal/ports"
	"github.com/alejandroik/trazavino/internal/service"
	"github.com/alejandroik/trazavino/pkg/server"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	app, shutdown := service.NewApplication(ctx)
	defer shutdown()

	server.RunHTTPServer(func(router *gin.Engine, prefix string) *gin.Engine {
		return ports.RegisterHandlersWithOptions(router, ports.NewHttpServer(app), ports.GinServerOptions{BaseURL: prefix})
	})
}
