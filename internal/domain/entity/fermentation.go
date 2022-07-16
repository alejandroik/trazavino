package entity

import "time"

type Fermentation struct {
	uuid string

	time time.Time

	warehouseUUID string
	warehouseName string

	tankUUID string
	tankName string

	hash        string
	transaction string
}
