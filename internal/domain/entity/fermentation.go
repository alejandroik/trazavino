package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Fermentation struct {
	Process

	warehouseUUID string
	warehouseName string

	tankUUID string
	tankName string
}

func NewFermentation(
	uuid string,
	startTime time.Time,
	wineryUUID string,
	warehouseUUID string,
	tankUUID string,
) (*Fermentation, error) {
	if uuid == "" {
		return nil, errors.New("empty Fermentation uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero Fermentation time")
	}
	if wineryUUID == "" {
		return nil, errors.New("empty winery uuid")
	}
	if warehouseUUID == "" {
		return nil, errors.New("empty warehouse uuid")
	}
	if tankUUID == "" {
		return nil, errors.New("empty tank uuid")
	}

	return &Fermentation{
		Process: Process{
			uuid:       uuid,
			wineryUUID: wineryUUID,
			startTime:  startTime,
		},
		warehouseUUID: warehouseUUID,
		tankUUID:      tankUUID,
	}, nil
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

func UnmarshalFermentationFromDatabase(
	uuid string,
	startTime time.Time,
	wineryUUID string,
	wineryName string,
	warehouseUUID string,
	warehouseName string,
	tankUUID string,
	tankName string,
	endTime time.Time,
	previousUUID string,
	hash string,
	transaction string,
) (*Fermentation, error) {
	f, err := NewFermentation(uuid, startTime, wineryUUID, warehouseUUID, tankUUID)
	if err != nil {
		return nil, err
	}

	f.wineryName = wineryName
	f.warehouseName = warehouseName
	f.tankName = tankName

	f.endTime = endTime
	f.previousUUID = previousUUID
	f.hash = hash
	f.transaction = transaction

	return f, nil
}
