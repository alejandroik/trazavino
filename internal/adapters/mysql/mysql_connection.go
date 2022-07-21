package mysql

import (
	"fmt"
	"os"

	gosql "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnection() (*gorm.DB, error) {
	cfg := &gosql.Config{
		Addr:   fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName: os.Getenv("DB_NAME"),
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
	}
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to MySQL")
	}

	return db, nil
}
