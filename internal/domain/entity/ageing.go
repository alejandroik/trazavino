package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Ageing struct {
	Process

	tankUUID string
	tankName string

	caskUUID string
	caskName string
}

func NewAgeing(
	uuid string,
	startTime time.Time,
	wineryUUID string,
	tankUUID string,
	caskUUID string,
) (*Ageing, error) {
	if uuid == "" {
		return nil, errors.New("empty Ageing uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero Ageing time")
	}
	if wineryUUID == "" {
		return nil, errors.New("empty winery uuid")
	}
	if tankUUID == "" {
		return nil, errors.New("empty tank uuid")
	}
	if caskUUID == "" {
		return nil, errors.New("empty cask uuid")
	}

	return &Ageing{
		Process: Process{
			uuid:       uuid,
			wineryUUID: wineryUUID,
			startTime:  startTime,
		},
		tankUUID: tankUUID,
		caskUUID: caskUUID,
	}, nil
}

func (a Ageing) TankUUID() string {
	return a.tankUUID
}

func (a Ageing) TankName() string {
	return a.tankName
}

func (a Ageing) CaskUUID() string {
	return a.caskUUID
}

func (a Ageing) CaskName() string {
	return a.caskName
}

func UnmarshalAgeingFromDatabase(
	uuid string,
	startTime time.Time,
	wineryUUID string,
	wineryName string,
	tankUUID string,
	tankName string,
	caskUUID string,
	caskName string,
	endTime time.Time,
	previousUUID string,
	hash string,
	transaction string,
) (*Ageing, error) {
	a, err := NewAgeing(uuid, startTime, wineryUUID, tankUUID, caskUUID)
	if err != nil {
		return nil, err
	}

	a.wineryName = wineryName
	a.tankName = tankName
	a.caskName = caskName

	a.endTime = endTime
	a.previousUUID = previousUUID
	a.hash = hash
	a.transaction = transaction

	return a, nil
}
