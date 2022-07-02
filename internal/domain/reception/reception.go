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

//TODO implement NewReception
func NewReception() (*Reception, error) {
	return &Reception{}, nil
}

func UnmarshallReceptionFromDatabase() (*Reception, error) {
	r, err := NewReception()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r Reception) Vineyard() vy.Vineyard {
	return r.vineyard
}

func (r Reception) Weight() int64 {
	return r.weight
}

func (r Reception) Sugar() int64 {
	return r.sugar
}
