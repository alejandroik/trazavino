package process

import (
	"time"
)

type Process struct {
	ID          int64
	StartDate   *time.Time
	EndDate     *time.Time
	Hash        string
	Transaction string
	Type        string
	Temperature int64
}
