package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Maceration struct {
	Process

	receptionUUID      string
	receptionStartTime time.Time

	warehouseUUID string
	warehouseName string
}

func NewMaceration(
	uuid string,
	startTime time.Time,
	wineryUUID string,
	receptionUUID string,
	warehouseUUID string,
) (*Maceration, error) {
	if uuid == "" {
		return nil, errors.New("empty maceration uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero maceration time")
	}
	if wineryUUID == "" {
		return nil, errors.New("empty winery uuid")
	}
	if receptionUUID == "" {
		return nil, errors.New("empty reception uuid")
	}
	if warehouseUUID == "" {
		return nil, errors.New("empty warehouse uuid")
	}

	return &Maceration{
		Process: Process{
			uuid:       uuid,
			wineryUUID: wineryUUID,
			startTime:  startTime,
		},
		receptionUUID: receptionUUID,
		warehouseUUID: warehouseUUID,
	}, nil
}

func (m Maceration) ReceptionUUID() string {
	return m.receptionUUID
}

func (m Maceration) ReceptionStartTime() time.Time {
	return m.receptionStartTime
}

func (m Maceration) WarehouseUUID() string {
	return m.warehouseUUID
}

func (m Maceration) WarehouseName() string {
	return m.warehouseName
}

func UnmarshalMacerationFromDatabase(
	uuid string,
	startTime time.Time,
	wineryUUID string,
	wineryName string,
	receptionUUID string,
	receptionStartTime time.Time,
	warehouseUUID string,
	warehouseName string,
	endTime time.Time,
	previousUUID string,
	hash string,
	transaction string) (*Maceration, error) {
	m, err := NewMaceration(uuid, startTime, wineryUUID, receptionUUID, warehouseUUID)
	if err != nil {
		return nil, err
	}

	m.wineryName = wineryName
	m.receptionStartTime = receptionStartTime
	m.warehouseName = warehouseName

	m.endTime = endTime
	m.previousUUID = previousUUID
	m.hash = hash
	m.transaction = transaction

	return m, nil
}
