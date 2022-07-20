package query

import "time"

type Reception struct {
	UUID string

	StartTime time.Time

	TruckUUID string
	Truck     string

	VineyardUUID string
	Vineyard     string

	GrapeTypeUUID string
	GrapeType     string

	Weight int32
	Sugar  int32

	EndTime     *time.Time
	Hash        *string
	Transaction *string
}
