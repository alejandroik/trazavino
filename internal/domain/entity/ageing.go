package entity

import "time"

type Ageing struct {
	uuid string

	startTime time.Time

	tankUUID string
	tankName string

	cellarUUID string
	cellarName string

	humidity int32

	endTime      time.Time
	previousUUID string

	hash        string
	transaction string
}
