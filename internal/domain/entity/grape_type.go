package entity

import "github.com/pkg/errors"

type GrapeType struct {
	ownedEntity
}

func NewGrapeType(id string, name string, wineryUUID string) (*GrapeType, error) {
	if id == "" {
		return nil, errors.New("empty uuid")
	}
	if name == "" {
		return nil, errors.New("empty name")
	}
	if wineryUUID == "" {
		return nil, errors.New("empty winery uuid")
	}

	return &GrapeType{ownedEntity{baseEntity{uuid: id, name: name}, wineryUUID}}, nil
}

func UnmarshalGrapeTypeFromDatabase(id string, wineryUUID string, name string) (*GrapeType, error) {
	return NewGrapeType(id, name, wineryUUID)
}
