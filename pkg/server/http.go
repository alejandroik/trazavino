package server

import (
	"os"

	"github.com/alejandroik/trazavino/internal/app"
	v1 "github.com/alejandroik/trazavino/internal/ports/v1"
	"github.com/alejandroik/trazavino/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RunHTTPServer(app app.Application, log logger.Interface) {
	RunHTTPServerOnAddr(":"+os.Getenv("SERVER_PORT"), app, log)
}

func RunHTTPServerOnAddr(addr string, app app.Application, log logger.Interface) {
	router := gin.Default()

	v1.RegisterHandlersWithOptions(router, v1.NewHttpServer(app), v1.GinServerOptions{BaseURL: "/v1"})

	log.Infof("Starting HTTP server on %s", addr)
	err := router.Run(addr)
	if err != nil {
		log.Fatalf("Unable to start HTTP server: %v", err)
	}
}
