package fermentation

import (
	"github.com/alejandroik/trazavino-api/internal/domain/process"
	"github.com/alejandroik/trazavino-api/internal/domain/vineyard"
)

type Fermentation struct {
	process   process.Process
	warehouse vineyard.Warehouse
	tank      vineyard.Tank
}
