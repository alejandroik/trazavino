package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Bottling struct {
	uuid string

	startTime time.Time

	caskUUID string
	caskName string

	wineUUID string
	wineName string

	bottleQty int32

	endTime      time.Time
	previousUUID string

	hash        string
	transaction string
}

func NewBottling(
	uuid string,
	startTime time.Time,
	caskUUID string,
	wineUUID string,
	bottleQty int32,
) (*Bottling, error) {
	if uuid == "" {
		return nil, errors.New("empty Bottling uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero Bottling time")
	}
	if caskUUID == "" {
		return nil, errors.New("empty cask uuid")
	}
	if wineUUID == "" {
		return nil, errors.New("empty wine uuid")
	}
	if bottleQty == 0 {
		return nil, errors.New("empty bottle quantity")
	}

	return &Bottling{
		uuid:      uuid,
		startTime: startTime,
		caskUUID:  caskUUID,
		wineUUID:  wineUUID,
		bottleQty: bottleQty,
	}, nil
}

func (a Bottling) UUID() string {
	return a.uuid
}

func (a Bottling) StartTime() time.Time {
	return a.startTime
}

func (a Bottling) CaskUUID() string {
	return a.caskUUID
}

func (a Bottling) CaskName() string {
	return a.caskName
}

func (a Bottling) WineUUID() string {
	return a.wineUUID
}

func (a Bottling) WineName() string {
	return a.wineName
}

func (a Bottling) BottleQty() int32 {
	return a.bottleQty
}

func (a Bottling) EndTime() time.Time {
	return a.endTime
}

func (a Bottling) PreviousUUID() string {
	return a.previousUUID
}

func (a Bottling) Hash() string {
	return a.hash
}

func (a Bottling) Transaction() string {
	return a.transaction
}

func (a *Bottling) UpdatePreviousUUID(pv string) error {
	a.previousUUID = pv

	return nil
}

func (a *Bottling) UpdateEndTime(t time.Time) error {
	a.endTime = t

	return nil
}

func (a *Bottling) UpdateHash(hash string) error {
	a.hash = hash

	return nil
}

func (a *Bottling) UpdateTransaction(tr string) error {
	a.transaction = tr

	return nil
}

func UnmarshalBottlingFromDatabase(
	uuid string,
	startTime time.Time,
	caskUUID string,
	caskName string,
	wineUUID string,
	wineName string,
	bottleQty int32,
	endTime time.Time,
	previousUUID string,
	hash string,
	transaction string,
) (*Bottling, error) {
	b, err := NewBottling(uuid, startTime, caskUUID, wineUUID, bottleQty)
	if err != nil {
		return nil, err
	}

	b.caskName = caskName
	b.wineName = wineName

	b.endTime = endTime
	b.previousUUID = previousUUID
	b.hash = hash
	b.transaction = transaction

	return b, nil
}
