package entity

type Truck struct {
	baseEntity
}

func NewTruck(id int64, name string) (*Truck, error) {
	return &Truck{baseEntity{id: id, name: name}}, nil
}

func UnmarshalTruckFromDatabase(id int64, name string) (*Truck, error) {
	return NewTruck(id, name)
}

func (t Truck) ID() int64 {
	return t.id
}

func (t Truck) Name() string {
	return t.name
}
