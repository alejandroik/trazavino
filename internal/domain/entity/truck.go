package entity

import "github.com/pkg/errors"

type Truck struct {
	ownedEntity
}

func NewTruck(id string, name string, wineryUUID string) (*Truck, error) {
	if id == "" {
		return nil, errors.New("empty uuid")
	}
	if name == "" {
		return nil, errors.New("empty name")
	}
	if wineryUUID == "" {
		return nil, errors.New("empty winery uuid")
	}

	return &Truck{ownedEntity{baseEntity{uuid: id, name: name}, wineryUUID}}, nil
}

func UnmarshalTruckFromDatabase(id string, name string, wineryUUID string) (*Truck, error) {
	return NewTruck(id, name, wineryUUID)
}
