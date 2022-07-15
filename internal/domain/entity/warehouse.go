package entity

type Warehouse struct {
	baseEntity
	isEmpty bool
}

func NewWarehouse(id int64, name string, isEmpty bool) (*Warehouse, error) {
	return &Warehouse{
		baseEntity: baseEntity{id: id, name: name},
		isEmpty:    isEmpty,
	}, nil
}

func (w Warehouse) ID() int64 {
	return w.id
}

func (w Warehouse) Name() string {
	return w.name
}

func (w Warehouse) IsEmpty() bool {
	return w.isEmpty
}

func UnmarshalWarehouseFromDatabase(id int64, name string, isEmpty bool) (*Warehouse, error) {
	return NewWarehouse(id, name, isEmpty)
}
