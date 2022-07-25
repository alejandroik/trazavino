package entity

import "github.com/pkg/errors"

type Wine struct {
	ownedEntity
}

func NewWine(id string, name string, wineryUUID string) (*Wine, error) {
	if id == "" {
		return nil, errors.New("empty uuid")
	}
	if name == "" {
		return nil, errors.New("empty name")
	}
	if wineryUUID == "" {
		return nil, errors.New("empty winery uuid")
	}

	return &Wine{ownedEntity{baseEntity{uuid: id, name: name}, wineryUUID}}, nil
}

func UnmarshalWineFromDatabase(id string, name string, wineryUUID string) (*Wine, error) {
	return NewWine(id, name, wineryUUID)
}
