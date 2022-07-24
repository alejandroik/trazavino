package entity

type Truck struct {
	baseEntity
	wineryUUID string
}

func NewTruck(id string, name string, wineryUUID string) (*Truck, error) {
	return &Truck{baseEntity{uuid: id, name: name}, wineryUUID}, nil
}

func UnmarshalTruckFromDatabase(id string, name string, wineryUUID string) (*Truck, error) {
	return NewTruck(id, name, wineryUUID)
}

func (t Truck) ID() string {
	return t.uuid
}

func (t Truck) WineryUUID() string {
	return t.wineryUUID
}

func (t Truck) Name() string {
	return t.name
}
