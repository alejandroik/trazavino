package entity

type Truck struct {
	baseEntity
}

func NewTruck(name string) (*Truck, error) {
	return newTruck(0, name)
}

func newTruck(id int, name string) (*Truck, error) {
	return &Truck{baseEntity{id: id, name: name}}, nil
}

func UnmarshalTruckFromDatabase(id int, name string) (*Truck, error) {
	return newTruck(id, name)
}

func (t Truck) ID() int {
	return t.id
}

func (t Truck) Name() string {
	return t.name
}
