package reception

import (
	"github.com/alejandroik/trazavino-api/internal/domain/process"
	vy "github.com/alejandroik/trazavino-api/internal/domain/vineyard"
)

type Reception struct {
	process   process.Process
	truck     vy.Truck
	vineyard  vy.Vineyard
	grapeType vy.GrapeType
	weight    int64
	sugar     int64
}
