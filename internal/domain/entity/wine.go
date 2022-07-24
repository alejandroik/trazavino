package entity

type Wine struct {
	baseEntity
}

func NewWine(id string, name string) (*Wine, error) {
	return &Wine{baseEntity{uuid: id, name: name}}, nil
}

func (w Wine) ID() string {
	return w.uuid
}

func (w Wine) Name() string {
	return w.name
}

func UnmarshalWineFromDatabase(id string, name string) (*Wine, error) {
	return NewWine(id, name)
}
