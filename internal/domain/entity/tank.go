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

func (w Tank) ID() string {
	return w.uuid
}

func (w Tank) Name() string {
	return w.name
}

func (w Tank) IsEmpty() bool {
	return w.isEmpty
}

func (w *Tank) UpdateIsEmpty(v bool) error {
	w.isEmpty = v

	return nil
}

func UnmarshalTankFromDatabase(id string, name string, isEmpty bool) (*Tank, error) {
	return NewTank(id, name, isEmpty)
}
