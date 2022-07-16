package ports

import (
	"net/http"

	"github.com/alejandroik/trazavino-api/internal/app"
	"github.com/alejandroik/trazavino-api/internal/app/command"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

// TODO implement
func (h HttpServer) RegisterReception(w http.ResponseWriter, r http.Request) {
	//postReception := PostReception{}

	cmd := command.RegisterReception{
		Weight:      0,
		Sugar:       0,
		TruckID:     0,
		VineyardID:  0,
		GrapeTypeID: 0,
	}

	err := h.app.Commands.RegisterReception.Handle(r.Context(), cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
