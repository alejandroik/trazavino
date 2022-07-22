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
	warehouseName string,
	tankUUID string,
	tankName string,
) (*Fermentation, error) {
	if uuid == "" {
		return nil, errors.New("empty maceration uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero maceration time")
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
		warehouseName: warehouseName,
		tankUUID:      tankUUID,
		tankName:      tankName,
	}, nil
}
