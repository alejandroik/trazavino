package maceration

import (
	"github.com/alejandroik/trazavino-api/internal/domain/process"
	"github.com/alejandroik/trazavino-api/internal/domain/reception"
	"github.com/alejandroik/trazavino-api/internal/domain/vineyard"
)

type Maceration struct {
	process   process.Process
	reception reception.Reception
	warehouse vineyard.Warehouse
}
