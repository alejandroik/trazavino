package entity

import "time"

type Ageing struct {
	uuid string

	time time.Time

	tankUUID string
	tankName string

	cellarUUID string
	cellarName string

	humidity int32

	hash        string
	transaction string
}
