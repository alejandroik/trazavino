package entity

type Truck struct {
	baseEntity
}

func NewTruck(id string, name string) (*Truck, error) {
	return &Truck{baseEntity{uuid: id, name: name}}, nil
}

func UnmarshalTruckFromDatabase(id string, name string) (*Truck, error) {
	return NewTruck(id, name)
}

func (t Truck) ID() string {
	return t.uuid
}

func (t Truck) Name() string {
	return t.name
}
