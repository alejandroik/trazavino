package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Fermentation struct {
	uuid string

	startTime time.Time

	warehouseUUID string
	warehouseName string

	tankUUID string
	tankName string

	endTime      time.Time
	previousUUID string

	hash        string
	transaction string
}

func NewFermentation(
	uuid string,
	startTime time.Time,
	warehouseUUID string,
	tankUUID string,
) (*Fermentation, error) {
	if uuid == "" {
		return nil, errors.New("empty Fermentation uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero Fermentation time")
	}
	if warehouseUUID == "" {
		return nil, errors.New("empty warehouse uuid")
	}
	if tankUUID == "" {
		return nil, errors.New("empty tank uuid")
	}

	return &Fermentation{
		uuid:          uuid,
		startTime:     startTime,
		warehouseUUID: warehouseUUID,
		tankUUID:      tankUUID,
	}, nil
}

func (f Fermentation) UUID() string {
	return f.uuid
}

func (f Fermentation) StartTime() time.Time {
	return f.startTime
}

func (f Fermentation) WarehouseUUID() string {
	return f.warehouseUUID
}

func (f Fermentation) WarehouseName() string {
	return f.warehouseName
}

func (f Fermentation) TankUUID() string {
	return f.tankUUID
}

func (f Fermentation) TankName() string {
	return f.tankName
}

func (f Fermentation) EndTime() time.Time {
	return f.endTime
}

func (f Fermentation) PreviousUUID() string {
	return f.previousUUID
}

func (f Fermentation) Hash() string {
	return f.hash
}

func (f Fermentation) Transaction() string {
	return f.transaction
}

func (f *Fermentation) UpdatePreviousUUID(pv string) error {
	f.previousUUID = pv

	return nil
}

func (f *Fermentation) UpdateEndTime(t time.Time) error {
	f.endTime = t

	return nil
}

func (f *Fermentation) UpdateHash(hash string) error {
	f.hash = hash

	return nil
}

func (f *Fermentation) UpdateTransaction(tr string) error {
	f.transaction = tr

	return nil
}

func UnmarshalFermentationFromDatabase(
	uuid string,
	startTime time.Time,
	warehouseUUID string,
	warehouseName string,
	tankUUID string,
	tankName string,
	endTime time.Time,
	previousUUID string,
	hash string,
	transaction string,
) (*Fermentation, error) {
	f, err := NewFermentation(uuid, startTime, warehouseUUID, tankUUID)
	if err != nil {
		return nil, err
	}

	f.warehouseName = warehouseName
	f.tankName = tankName

	f.endTime = endTime
	f.previousUUID = previousUUID
	f.hash = hash
	f.transaction = transaction

	return f, nil
}
