package entity

import "github.com/pkg/errors"

type Winery struct {
	baseEntity
}

func NewWinery(id string, name string) (*Winery, error) {
	if id == "" {
		return nil, errors.New("empty uuid")
	}
	if name == "" {
		return nil, errors.New("empty name")
	}

	return &Winery{baseEntity{uuid: id, name: name}}, nil
}

func UnmarshalWineryFromDatabase(id string, name string) (*Winery, error) {
	return NewWinery(id, name)
}
