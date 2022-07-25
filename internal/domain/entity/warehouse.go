package entity

type Warehouse struct {
	ownedEntity
	isEmpty bool
}

func NewWarehouse(id string, name string, isEmpty bool, wineryUUID string) (*Warehouse, error) {
	return &Warehouse{
		ownedEntity: ownedEntity{baseEntity{uuid: id, name: name}, wineryUUID},
		isEmpty:     isEmpty,
	}, nil
}

func (w Warehouse) IsEmpty() bool {
	return w.isEmpty
}

func (w *Warehouse) UpdateIsEmpty(v bool) error {
	w.isEmpty = v

	return nil
}

func UnmarshalWarehouseFromDatabase(id string, name string, isEmpty bool, wineryUUID string) (*Warehouse, error) {
	return NewWarehouse(id, name, isEmpty, wineryUUID)
}
