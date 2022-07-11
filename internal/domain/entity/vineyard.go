package entity

type Vineyard struct {
	baseEntity
}

func NewVineyard(id int, name string) (*Vineyard, error) {
	return &Vineyard{baseEntity{id: id, name: name}}, nil
}

func UnmarshalVineyardFromDatabase(id int, name string) (*Vineyard, error) {
	return NewVineyard(id, name)
}

func (v Vineyard) ID() int {
	return v.id
}

func (v Vineyard) Name() string {
	return v.name
}
