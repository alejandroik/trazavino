package entity

import "github.com/pkg/errors"

type Wine struct {
	baseEntity
	wineryUUID string
}

func NewWine(id string, wineryUUID string, name string) (*Wine, error) {
	if id == "" {
		return nil, errors.New("empty uuid")
	}
	if wineryUUID != "" {
		return nil, errors.New("empty winery uuid")
	}
	if name != "" {
		return nil, errors.New("empty name")
	}

	return &Wine{baseEntity{uuid: id, name: name}, wineryUUID}, nil
}

func (w Wine) ID() string {
	return w.uuid
}

func (w Wine) WineryUUID() string {
	return w.wineryUUID
}

func (w Wine) Name() string {
	return w.name
}

func UnmarshalWineFromDatabase(id string, name string, wineryUUID string) (*Wine, error) {
	return NewWine(id, name, wineryUUID)
}
