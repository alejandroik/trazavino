package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Maceration struct {
	uuid string

	time time.Time

	receptionUUID string

	warehouseUUID string
	warehouseName string

	hash        string
	transaction string
}

func NewMaceration(uuid string, time time.Time, receptionUUID string, warehouseUUID string, warehouseName string) (*Maceration, error) {
	if uuid == "" {
		return nil, errors.New("empty maceration uuid")
	}
	if time.IsZero() {
		return nil, errors.New("zero maceration time")
	}
	if receptionUUID == "" {
		return nil, errors.New("empty reception uuid")
	}
	if warehouseUUID == "" {
		return nil, errors.New("empty warehouse uuid")
	}
	if warehouseName == "" {
		return nil, errors.New("empty warehouse name")
	}

	return &Maceration{
		uuid:          uuid,
		time:          time,
		receptionUUID: receptionUUID,
		warehouseUUID: warehouseUUID,
		warehouseName: warehouseName,
	}, nil
}

func (m Maceration) UUID() string {
	return m.uuid
}

func (m Maceration) ReceptionUUID() string {
	return m.receptionUUID
}

func (m Maceration) WarehouseUUID() string {
	return m.warehouseUUID
}

func (m Maceration) WarehouseName() string {
	return m.warehouseName
}

func (r *Maceration) UpdateHash(hash string) error {
	r.hash = hash

	return nil
}

func (r *Maceration) UpdateTransaction(tr string) error {
	r.transaction = tr

	return nil
}

func UnmarshalMacerationFromDatabase(
	uuid string,
	time time.Time,
	receptionUUID string,
	warehouseUUID string,
	warehouseName string,
	hash string,
	transaction string) (*Maceration, error) {
	m, err := NewMaceration(uuid, time, receptionUUID, warehouseUUID, warehouseName)
	if err != nil {
		return nil, err
	}

	m.hash = hash
	m.transaction = transaction

	return m, nil
}
