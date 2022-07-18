package ports

import (
	"net/http"
	"time"

	"github.com/alejandroik/trazavino/internal/app"
	"github.com/alejandroik/trazavino/internal/app/command"
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
		ReceptionUUID:      "",
		ReceptionStartTime: time.Time{},
		TruckUUID:          "",
		TruckLicense:       "",
		VineyardUUID:       "",
		VineyardName:       "",
		GrapeTypeUUID:      "",
		GrapeTypeName:      "",
		Weight:             0,
		Sugar:              0,
	}

	err := h.app.Commands.RegisterReception.Handle(r.Context(), cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
