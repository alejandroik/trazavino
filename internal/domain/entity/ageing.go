package entity

import (
	"github.com/pkg/errors"
	"time"
)

type Ageing struct {
	uuid string

	startTime time.Time

	tankUUID string
	tankName string

	caskUUID string
	caskName string

	humidity int32

	endTime      time.Time
	previousUUID string

	hash        string
	transaction string
}

func NewAgeing(
	uuid string,
	startTime time.Time,
	tankUUID string,
	caskUUID string,
) (*Ageing, error) {
	if uuid == "" {
		return nil, errors.New("empty Ageing uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero Ageing time")
	}
	if tankUUID == "" {
		return nil, errors.New("empty tank uuid")
	}
	if caskUUID == "" {
		return nil, errors.New("empty cask uuid")
	}

	return &Ageing{
		uuid:      uuid,
		startTime: startTime,
		tankUUID:  tankUUID,
		caskUUID:  caskUUID,
	}, nil
}

func (a Ageing) UUID() string {
	return a.uuid
}

func (a Ageing) StartTime() time.Time {
	return a.startTime
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

func (a Ageing) EndTime() time.Time {
	return a.endTime
}

func (a Ageing) PreviousUUID() string {
	return a.previousUUID
}

func (a Ageing) Hash() string {
	return a.hash
}

func (a Ageing) Transaction() string {
	return a.transaction
}

func (a *Ageing) UpdatePreviousUUID(pv string) error {
	a.previousUUID = pv

	return nil
}

func (a *Ageing) UpdateEndTime(t time.Time) error {
	a.endTime = t

	return nil
}

func (a *Ageing) UpdateHash(hash string) error {
	a.hash = hash

	return nil
}

func (a *Ageing) UpdateTransaction(tr string) error {
	a.transaction = tr

	return nil
}

func UnmarshalAgeingFromDatabase(
	uuid string,
	startTime time.Time,
	tankUUID string,
	tankName string,
	caskUUID string,
	caskName string,
	endTime time.Time,
	previousUUID string,
	hash string,
	transaction string,
) (*Ageing, error) {
	f, err := NewAgeing(uuid, startTime, tankUUID, caskUUID)
	if err != nil {
		return nil, err
	}

	f.tankName = tankName
	f.caskName = caskName

	f.endTime = endTime
	f.previousUUID = previousUUID
	f.hash = hash
	f.transaction = transaction

	return f, nil
}
