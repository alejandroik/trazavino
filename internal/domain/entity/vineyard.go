package entity

type Vineyard struct {
	baseEntity
}

func NewVineyard(id string, name string) (*Vineyard, error) {
	return &Vineyard{baseEntity{uuid: id, name: name}}, nil
}

func UnmarshalVineyardFromDatabase(id string, name string) (*Vineyard, error) {
	return NewVineyard(id, name)
}

func (v Vineyard) ID() string {
	return v.uuid
}

func (v Vineyard) Name() string {
	return v.name
}
