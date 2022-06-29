package ageing

import (
	"github.com/alejandroik/trazavino-api/internal/domain/process"
	"github.com/alejandroik/trazavino-api/internal/domain/vineyard"
)

type Ageing struct {
	process  process.Process
	tank     vineyard.Tank
	cellar   vineyard.Cellar
	humidity int64
}
