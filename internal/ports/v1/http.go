package v1

import "github.com/alejandroik/trazavino/internal/app"

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}
