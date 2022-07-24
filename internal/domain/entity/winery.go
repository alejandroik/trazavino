package entity

type Winery struct {
	baseEntity
}

func NewWinery(id string, name string) (*Winery, error) {
	return &Winery{baseEntity{uuid: id, name: name}}, nil
}

func (w Winery) ID() string {
	return w.uuid
}

func (w Winery) Name() string {
	return w.name
}

func UnmarshalWineryFromDatabase(id string, name string) (*Winery, error) {
	return NewWinery(id, name)
}
