package entity

import "github.com/pkg/errors"

type Vineyard struct {
	ownedEntity
}

func NewVineyard(id string, name string, wineryUUID string) (*Vineyard, error) {
	if id == "" {
		return nil, errors.New("empty uuid")
	}
	if name == "" {
		return nil, errors.New("empty name")
	}
	if wineryUUID == "" {
		return nil, errors.New("empty winery uuid")
	}

	return &Vineyard{ownedEntity{baseEntity{uuid: id, name: name}, wineryUUID}}, nil
}

func UnmarshalVineyardFromDatabase(id string, name string, wineryUUID string) (*Vineyard, error) {
	return NewVineyard(id, name, wineryUUID)
}
