package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func RunHTTPServer(createHandler func(router *gin.Engine, prefix string) *gin.Engine) {
	RunHTTPServerOnAddr(":"+os.Getenv("PORT"), createHandler)
}

func RunHTTPServerOnAddr(addr string, createHandler func(router *gin.Engine, prefix string) *gin.Engine) {
	apiRouter := gin.Default()

	createHandler(apiRouter, "/v1")

	log.Println("Starting HTTP server")

	err := apiRouter.Run(addr)
	if err != nil {
		log.Fatalf("Unable to start HTTP server: %v", err)
	}
}
