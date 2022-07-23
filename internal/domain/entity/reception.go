package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Reception struct {
	uuid string

	startTime time.Time

	truckUUID    string
	truckLicense string

	vineyardUUID string
	vineyardName string

	grapeTypeUUID string
	grapeTypeName string

	weight int32
	sugar  int32

	endTime     time.Time
	hash        string
	transaction string
}

func NewReception(
	uuid string,
	startTime time.Time,
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
	if truckUUID == "" {
		return nil, errors.New("empty truck uuid")
	}
	if vineyardUUID == "" {
		return nil, errors.New("empty vineyard uuid")
	}
	if grapeTypeUUID == "" {
		return nil, errors.New("empty grape type uuid")
	}

	return &Reception{
		uuid:          uuid,
		startTime:     startTime,
		truckUUID:     truckUUID,
		vineyardUUID:  vineyardUUID,
		grapeTypeUUID: grapeTypeUUID,
		weight:        weight,
		sugar:         sugar,
	}, nil
}

func (r Reception) UUID() string {
	return r.uuid
}

func (r Reception) StartTime() time.Time {
	return r.startTime
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

func (r Reception) EndTime() time.Time {
	return r.endTime
}

func (r Reception) Hash() string {
	return r.hash
}

func (r Reception) Transaction() string {
	return r.transaction
}

func (r *Reception) UpdateEndTime(t time.Time) error {
	r.endTime = t

	return nil
}

func (r *Reception) UpdateHash(hash string) error {
	r.hash = hash

	return nil
}

func (r *Reception) UpdateTransaction(tr string) error {
	r.transaction = tr

	return nil
}

func UnmarshalReceptionFromDatabase(
	uuid string,
	startTime time.Time,
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
		truckUUID,
		vineyardUUID,
		grapeTypeUUID,
		weight,
		sugar)
	if err != nil {
		return nil, err
	}

	r.truckLicense = truckLicense
	r.vineyardName = vineyardName
	r.grapeTypeName = grapeTypeName

	r.endTime = endTime
	r.hash = hash
	r.transaction = transaction

	return r, nil
}
