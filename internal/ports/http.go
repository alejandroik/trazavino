package ports

import (
	"net/http"

	"github.com/alejandroik/trazavino-api/internal/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) RegisterReception(w http.ResponseWriter, r http.Request) {

}
