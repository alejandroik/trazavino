package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Reception struct {
	Process

	truckUUID    string
	truckLicense string

	vineyardUUID string
	vineyardName string

	grapeTypeUUID string
	grapeTypeName string

	weight int32
	sugar  int32
}

func NewReception(
	uuid string,
	startTime time.Time,
	wineryUUID string,
	truckUUID string,
	vineyardUUID string,
	grapeTypeUUID string,
	weight int32,
	sugar int32) (*Reception, error) {
	if uuid == "" {
		return nil, errors.New("empty reception uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero reception start time")
	}
	if wineryUUID == "" {
		return nil, errors.New("empty winery uuid")
	}
	if truckUUID == "" {
		return nil, errors.New("empty truck uuid")
	}
	if vineyardUUID == "" {
		return nil, errors.New("empty vineyard uuid")
	}
	if grapeTypeUUID == "" {
		return nil, errors.New("empty grape type uuid")
	}
	if weight == 0 {
		return nil, errors.New("empty weight")
	}
	if sugar == 0 {
		return nil, errors.New("empty sugar")
	}

	return &Reception{
		Process: Process{
			uuid:       uuid,
			wineryUUID: wineryUUID,
			startTime:  startTime,
		},
		truckUUID:     truckUUID,
		vineyardUUID:  vineyardUUID,
		grapeTypeUUID: grapeTypeUUID,
		weight:        weight,
		sugar:         sugar,
	}, nil
}

func (r Reception) TruckUUID() string {
	return r.truckUUID
}

func (r Reception) TruckLicense() string {
	return r.truckLicense
}

func (r Reception) VineyardUUID() string {
	return r.vineyardUUID
}

func (r Reception) VineyardName() string {
	return r.vineyardName
}

func (r Reception) GrapeTypeUUID() string {
	return r.grapeTypeUUID
}

func (r Reception) GrapeTypeName() string {
	return r.grapeTypeName
}

func (r Reception) Weight() int32 {
	return r.weight
}

func (r Reception) Sugar() int32 {
	return r.sugar
}

func UnmarshalReceptionFromDatabase(
	uuid string,
	startTime time.Time,
	wineryUUID string,
	wineryName string,
	truckUUID string,
	truckLicense string,
	vineyardUUID string,
	vineyardName string,
	grapeTypeUUID string,
	grapeTypeName string,
	weight int32,
	sugar int32,
	endTime time.Time,
	hash string,
	transaction string) (*Reception, error) {
	r, err := NewReception(
		uuid,
		startTime,
		wineryUUID,
		truckUUID,
		vineyardUUID,
		grapeTypeUUID,
		weight,
		sugar)
	if err != nil {
		return nil, err
	}

	r.wineryName = wineryName
	r.truckLicense = truckLicense
	r.vineyardName = vineyardName
	r.grapeTypeName = grapeTypeName

	r.endTime = endTime
	r.hash = hash
	r.transaction = transaction

	return r, nil
}
