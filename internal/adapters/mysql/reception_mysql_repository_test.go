package mysql

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/alejandroik/trazavino-api/internal/domain/process"
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
		Passwd:    "",
		ParseTime: true,
	}
	mysqlDb, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	mysqlDb.AutoMigrate(&ProcessModel{}, &ReceptionModel{})

	rec, _ := reception.NewReception(process.Process{}, 5, 10)

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
			rm, err := r.AddReception(tt.args.ctx, tt.args.rc)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddReception() error = %v, wantErr %v", err, tt.wantErr)
			}
			rec, err := r.GetReception(context.Background(), int64(rm.ID))
			if err != nil {
				t.Error(err)
			}
			t.Log(rec)
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
