package ports

import (
	"github.com/alejandroik/trazavino-api/internal/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}
