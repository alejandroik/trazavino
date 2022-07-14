package entity

type Vineyard struct {
	baseEntity
}

func NewVineyard(id int64, name string) (*Vineyard, error) {
	return &Vineyard{baseEntity{id: id, name: name}}, nil
}

func UnmarshalVineyardFromDatabase(id int64, name string) (*Vineyard, error) {
	return NewVineyard(id, name)
}

func (v Vineyard) ID() int64 {
	return v.id
}

func (v Vineyard) Name() string {
	return v.name
}
