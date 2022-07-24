package entity

type Tank struct {
	baseEntity
	isEmpty bool
}

func NewTank(id string, name string, isEmpty bool) (*Tank, error) {
	return &Tank{
		baseEntity: baseEntity{uuid: id, name: name},
		isEmpty:    isEmpty,
	}, nil
}

func (t Tank) ID() string {
	return t.uuid
}

func (t Tank) Name() string {
	return t.name
}

func (t Tank) IsEmpty() bool {
	return t.isEmpty
}

func (t *Tank) UpdateIsEmpty(v bool) error {
	t.isEmpty = v

	return nil
}

func UnmarshalTankFromDatabase(id string, name string, isEmpty bool) (*Tank, error) {
	return NewTank(id, name, isEmpty)
}
