package adapters

import (
	"time"

	"gorm.io/gorm"
)

type ProcessModel struct {
	gorm.Model
	StartDate   *time.Time
	EndDate     *time.Time
	Hash        string
	Transaction string
	Type        string
	Temperature int64
}
