package entity

import "time"

type Bottling struct {
	uuid string

	startTime time.Time

	cellarUUID string
	cellarName string

	wineUUID string
	wineName string

	bottleQty int32

	endTime      time.Time
	previousUUID string

	hash        string
	transaction string
}
