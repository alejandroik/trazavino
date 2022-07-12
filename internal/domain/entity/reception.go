package entity

type Reception struct {
	process   *Process
	truck     *Truck
	vineyard  *Vineyard
	grapeType *GrapeType
	weight    int
	sugar     int
}

func NewReception(
	process *Process,
	truck *Truck,
	//vineyard vy.Vineyard,
	//grapeType vy.GrapeType,
	weight int,
	sugar int) (*Reception, error) {
	return newReception(
		process, truck, nil, nil, weight, sugar)
}

func newReception(
	process *Process,
	truck *Truck,
	vineyard *Vineyard,
	grapeType *GrapeType,
	weight int,
	sugar int) (*Reception, error) {
	return &Reception{
		process:   process,
		truck:     truck,
		vineyard:  vineyard,
		grapeType: grapeType,
		weight:    weight,
		sugar:     sugar,
	}, nil
}

func UnmarshalReceptionFromDatabase(
	process *Process,
	truck *Truck,
	//vineyard vy.Vineyard,
	//grapeType vy.GrapeType,
	weight int,
	sugar int) (*Reception, error) {
	r, err := newReception(process, truck, nil, nil, weight, sugar)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r Reception) Process() *Process {
	return r.process
}

func (r Reception) Truck() *Truck {
	return r.truck
}

func (r Reception) Vineyard() *Vineyard {
	return r.vineyard
}

func (r Reception) GrapeType() *GrapeType {
	return r.grapeType
}

func (r Reception) Weight() int {
	return r.weight
}

func (r Reception) Sugar() int {
	return r.sugar
}
