package mysql

import "gorm.io/gorm"

type BaseEntityModel struct {
	gorm.Model
	Name string
}
