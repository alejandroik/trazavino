package entity

type Tank struct {
	ownedEntity
	isEmpty bool
}

func NewTank(id string, name string, isEmpty bool, wineryUUID string) (*Tank, error) {
	return &Tank{
		ownedEntity: ownedEntity{baseEntity{uuid: id, name: name}, wineryUUID},
		isEmpty:     isEmpty,
	}, nil
}

func (t Tank) IsEmpty() bool {
	return t.isEmpty
}

func (t *Tank) UpdateIsEmpty(v bool) error {
	t.isEmpty = v

	return nil
}

func UnmarshalTankFromDatabase(id string, name string, isEmpty bool, wineryUUID string) (*Tank, error) {
	return NewTank(id, name, isEmpty, wineryUUID)
}
