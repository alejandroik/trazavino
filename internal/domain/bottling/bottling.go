package bottling

import (
	"github.com/alejandroik/trazavino-api/internal/domain/process"
	"github.com/alejandroik/trazavino-api/internal/domain/vineyard"
)

type Bottling struct {
	process   process.Process
	cellar    vineyard.Cellar
	bottleQty int64
	wine      vineyard.Wine
}
