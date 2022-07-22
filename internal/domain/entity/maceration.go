package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Maceration struct {
	uuid string

	startTime time.Time

	receptionUUID      string
	receptionStartTime time.Time

	warehouseUUID string
	warehouseName string

	endTime     time.Time
	hash        string
	transaction string
}

func NewMaceration(uuid string,
	startTime time.Time,
	receptionUUID string,
	receptionStartTime time.Time,
	warehouseUUID string,
	warehouseName string,
) (*Maceration, error) {
	if uuid == "" {
		return nil, errors.New("empty maceration uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero maceration time")
	}
	if receptionUUID == "" {
		return nil, errors.New("empty reception uuid")
	}
	if receptionStartTime.IsZero() {
		return nil, errors.New("zero reception time")
	}
	if warehouseUUID == "" {
		return nil, errors.New("empty warehouse uuid")
	}
	if warehouseName == "" {
		return nil, errors.New("empty warehouse name")
	}

	return &Maceration{
		uuid:               uuid,
		startTime:          startTime,
		receptionUUID:      receptionUUID,
		receptionStartTime: receptionStartTime,
		warehouseUUID:      warehouseUUID,
		warehouseName:      warehouseName,
	}, nil
}

func (m Maceration) UUID() string {
	return m.uuid
}

func (m Maceration) StartTime() time.Time {
	return m.startTime
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

func (m Maceration) EndTime() time.Time {
	return m.endTime
}

func (m Maceration) Hash() string {
	return m.hash
}

func (m Maceration) Transaction() string {
	return m.transaction
}

func (m *Maceration) UpdateEndTime(t time.Time) error {
	m.endTime = t

	return nil
}

func (m *Maceration) UpdateHash(hash string) error {
	m.hash = hash

	return nil
}

func (m *Maceration) UpdateTransaction(tr string) error {
	m.transaction = tr

	return nil
}

func UnmarshalMacerationFromDatabase(
	uuid string,
	time time.Time,
	receptionUUID string,
	receptionStartTime time.Time,
	warehouseUUID string,
	warehouseName string,
	endTime time.Time,
	hash string,
	transaction string) (*Maceration, error) {
	m, err := NewMaceration(uuid, time, receptionUUID, receptionStartTime, warehouseUUID, warehouseName)
	if err != nil {
		return nil, err
	}

	m.endTime = endTime
	m.hash = hash
	m.transaction = transaction

	return m, nil
}
