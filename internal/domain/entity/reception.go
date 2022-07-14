package entity

type Reception struct {
	id          int64
	truckId     int64
	vineyardId  int64
	grapeTypeId int64
	weight      int32
	sugar       int32
}

func NewReception(
	id int64,
	truckId int64,
	vineyardId int64,
	grapeTypeId int64,
	weight int32,
	sugar int32) (*Reception, error) {
	return &Reception{
		id:          id,
		truckId:     truckId,
		vineyardId:  vineyardId,
		grapeTypeId: grapeTypeId,
		weight:      weight,
		sugar:       sugar,
	}, nil
}

func (r Reception) ID() int64 {
	return r.id
}

func (r Reception) TruckID() int64 {
	return r.truckId
}

func (r Reception) VineyardID() int64 {
	return r.vineyardId
}

func (r Reception) GrapeTypeID() int64 {
	return r.grapeTypeId
}

func (r Reception) Weight() int32 {
	return r.weight
}

func (r Reception) Sugar() int32 {
	return r.sugar
}

func UnmarshalReceptionFromDatabase(
	id int64,
	truckId int64,
	vineyardId int64,
	grapeTypeId int64,
	weight int32,
	sugar int32) (*Reception, error) {
	r, err := NewReception(id, truckId, vineyardId, grapeTypeId, weight, sugar)
	if err != nil {
		return nil, err
	}

	return r, nil
}
