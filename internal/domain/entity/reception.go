package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Reception struct {
	uuid string

	time time.Time

	truckUUID    string
	truckLicense string

	vineyardUUID string
	vineyardName string

	grapeTypeUUID string
	grapeTypeName string

	weight int32
	sugar  int32

	hash        string
	transaction string
}

func NewReception(
	uuid string,
	time time.Time,
	truckUUID string,
	truckLicense string,
	vineyardUUID string,
	vineyardName string,
	grapeTypeUUID string,
	grapeTypeName string,
	weight int32,
	sugar int32) (*Reception, error) {
	if uuid == "" {
		return nil, errors.New("empty reception uuid")
	}
	if time.IsZero() {
		return nil, errors.New("zero reception start time")
	}
	if truckUUID == "" {
		return nil, errors.New("empty truck uuid")
	}
	if truckLicense == "" {
		return nil, errors.New("empty truck license")
	}
	if vineyardUUID == "" {
		return nil, errors.New("empty vineyard uuid")
	}
	if vineyardName == "" {
		return nil, errors.New("empty vineyard name")
	}
	if grapeTypeUUID == "" {
		return nil, errors.New("empty grape type uuid")
	}
	if grapeTypeName == "" {
		return nil, errors.New("empty grape type name")
	}

	return &Reception{
		uuid:          uuid,
		time:          time,
		truckUUID:     truckUUID,
		truckLicense:  truckLicense,
		vineyardUUID:  vineyardUUID,
		vineyardName:  vineyardName,
		grapeTypeUUID: grapeTypeUUID,
		grapeTypeName: grapeTypeName,
		weight:        weight,
		sugar:         sugar,
	}, nil
}

func (r Reception) UUID() string {
	return r.uuid
}

func (r Reception) Time() time.Time {
	return r.time
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
	time time.Time,
	truckUUID string,
	truckLicense string,
	vineyardUUID string,
	vineyardName string,
	grapeTypeUUID string,
	grapeTypeName string,
	weight int32,
	sugar int32,
	hash string,
	transaction string) (*Reception, error) {
	r, err := NewReception(
		uuid,
		time,
		truckUUID,
		truckLicense,
		vineyardUUID,
		vineyardName,
		grapeTypeUUID,
		grapeTypeName,
		weight,
		sugar)
	if err != nil {
		return nil, err
	}

	r.hash = hash
	r.transaction = transaction

	return r, nil
}
