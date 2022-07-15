package app

import "github.com/alejandroik/trazavino-api/internal/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	RegisterReception  command.RegisterReceptionHandler
	RegisterMaceration command.RegisterMacerationHandler
}

type Queries struct {
}
