package entity

type Maceration struct {
	id          int64
	receptionId int64
	warehouseId int64
}

func NewMaceration(id int64, receptionId int64, warehouseId int64) (*Maceration, error) {
	return &Maceration{
		id:          id,
		receptionId: receptionId,
		warehouseId: warehouseId,
	}, nil
}

func (m Maceration) ID() int64 {
	return m.id
}

func (m Maceration) ReceptionID() int64 {
	return m.receptionId
}

func (m Maceration) WarehouseID() int64 {
	return m.warehouseId
}

func UnmarshalMacerationFromDatabase(id int64, receptionId int64, warehouseId int64) (*Maceration, error) {
	return NewMaceration(id, receptionId, warehouseId)
}
