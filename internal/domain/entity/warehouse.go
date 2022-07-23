package entity

type Warehouse struct {
	baseEntity
	isEmpty bool
}

func NewWarehouse(id string, name string, isEmpty bool) (*Warehouse, error) {
	return &Warehouse{
		baseEntity: baseEntity{uuid: id, name: name},
		isEmpty:    isEmpty,
	}, nil
}

func (w Warehouse) ID() string {
	return w.uuid
}

func (w Warehouse) Name() string {
	return w.name
}

func (w Warehouse) IsEmpty() bool {
	return w.isEmpty
}

func (w *Warehouse) UpdateIsEmpty(v bool) error {
	w.isEmpty = v

	return nil
}

func UnmarshalWarehouseFromDatabase(id string, name string, isEmpty bool) (*Warehouse, error) {
	return NewWarehouse(id, name, isEmpty)
}
