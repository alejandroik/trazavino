package mysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	gosql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestReceptionMysqlRepository_AddReception(t *testing.T) {
	ctx := context.Background()
	cfg := &gosql.Config{
		Addr:      fmt.Sprintf("%v:%v", "localhost", "3306"),
		DBName:    "trazavino",
		User:      "root",
		Passwd:    "",
		ParseTime: true,
	}
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.Begin()
	db.AutoMigrate(&ProcessModel{}, &ReceptionModel{}, &TruckModel{})

	truck, _ := entity.NewTruck("EEOO-990")
	tr := NewTruckMysqlRepository(db)
	savedTruck, err := tr.AddTruck(ctx, truck)
	if err != nil {
		t.Error(err)
	}
	t.Log("Added truck: ", savedTruck)

	rec, _ := entity.NewReception(savedTruck, 5, 10)
	rr := NewReceptionMysqlRepository(db)
	savedRec, err := rr.AddReception(ctx, rec)
	if err != nil {
		t.Error(err)
	}
	t.Log("Added reception: ", savedRec)
	db.Rollback()
}
