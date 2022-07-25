package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Bottling struct {
	Process

	caskUUID string
	caskName string

	wineUUID string
	wineName string

	bottleQty int32
}

func NewBottling(
	uuid string,
	startTime time.Time,
	wineryUUID string,
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
	if wineryUUID == "" {
		return nil, errors.New("empty winery uuid")
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
		Process: Process{
			uuid:       uuid,
			wineryUUID: wineryUUID,
			startTime:  startTime,
		},
		caskUUID:  caskUUID,
		wineUUID:  wineUUID,
		bottleQty: bottleQty,
	}, nil
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

func UnmarshalBottlingFromDatabase(
	uuid string,
	startTime time.Time,
	wineryUUID string,
	wineryName string,
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
	b, err := NewBottling(uuid, startTime, wineryUUID, caskUUID, wineUUID, bottleQty)
	if err != nil {
		return nil, err
	}

	b.wineryName = wineryName
	b.caskName = caskName
	b.wineName = wineName

	b.endTime = endTime
	b.previousUUID = previousUUID
	b.hash = hash
	b.transaction = transaction

	return b, nil
}
