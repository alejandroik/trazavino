package entity

import "time"

type Bottling struct {
	uuid string

	time time.Time

	cellarUUID string
	cellarName string

	wineUUID string
	wineName string

	bottleQty int32

	hash        string
	transaction string
}
