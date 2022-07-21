package main

import (
	"context"

	"github.com/alejandroik/trazavino/internal/service"
	"github.com/alejandroik/trazavino/pkg/logger"
	"github.com/alejandroik/trazavino/pkg/server"
)

func main() {
	log := logger.New()

	ctx := context.Background()

	app := service.NewApplication(ctx, log)

	server.RunHTTPServer(app, log)
}
