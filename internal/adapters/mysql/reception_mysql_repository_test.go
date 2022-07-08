package mysql

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/alejandroik/trazavino-api/internal/domain/reception"
	gosql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestReceptionMysqlRepository_AddReception(t *testing.T) {
	cfg := &gosql.Config{
		Addr:      fmt.Sprintf("%v:%v", "localhost", "3306"),
		DBName:    "trazavino",
		User:      "root",
		Passwd:    "root",
		ParseTime: true,
	}
	mysqlDb, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	mysqlDb.AutoMigrate(&ProcessModel{}, &ReceptionModel{})

	rec, _ := reception.NewReception(5, 10)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		rc  *reception.Reception
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields:  fields{db: mysqlDb},
			wantErr: false,
			args: args{
				rc: rec,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReceptionMysqlRepository{
				db: tt.fields.db,
			}
			if err := r.AddReception(tt.args.ctx, tt.args.rc); (err != nil) != tt.wantErr {
				t.Errorf("AddReception() error = %v, wantErr %v", err, tt.wantErr)
			}
			rm := &ReceptionModel{}
			result := r.db.First(rm)
			if result.Error != nil {
				t.Error(err)
			}
			t.Log(rm)
			pm := &ProcessModel{}
			result = r.db.First(pm)
			if result.Error != nil {
				t.Error(err)
			}
			t.Log(pm)
		})
	}
}

func TestReceptionMysqlRepository_GetReception(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx         context.Context
		receptionId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *reception.Reception
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReceptionMysqlRepository{
				db: tt.fields.db,
			}
			got, err := r.GetReception(tt.args.ctx, tt.args.receptionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReception() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReception() got = %v, want %v", got, tt.want)
			}
		})
	}
}
