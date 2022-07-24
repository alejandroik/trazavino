package entity

type Bottle struct {
	baseEntity
}

func NewBottle(id string, name string) (*Bottle, error) {
	return &Bottle{baseEntity{uuid: id, name: name}}, nil
}

func (b Bottle) ID() string {
	return b.uuid
}

func (b Bottle) Name() string {
	return b.name
}

func UnmarshalBottleFromDatabase(id string, name string) (*Bottle, error) {
	return NewBottle(id, name)
}
