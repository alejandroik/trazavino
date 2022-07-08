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

func NewReception(
	//truck vy.Truck,
	//vineyard vy.Vineyard,
	//grapeType vy.GrapeType,
	weight int64,
	sugar int64) (*Reception, error) {
	return &Reception{
		weight: weight,
		sugar:  sugar}, nil
}

func UnmarshallReceptionFromDatabase(
	//truck vy.Truck,
	//vineyard vy.Vineyard,
	//grapeType vy.GrapeType,
	weight int64,
	sugar int64) (*Reception, error) {
	r, err := NewReception(weight, sugar)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r Reception) Truck() vy.Truck {
	return r.truck
}

func (r Reception) Vineyard() vy.Vineyard {
	return r.vineyard
}

func (r Reception) GrapeType() vy.GrapeType {
	return r.grapeType
}

func (r Reception) Weight() int64 {
	return r.weight
}

func (r Reception) Sugar() int64 {
	return r.sugar
}
