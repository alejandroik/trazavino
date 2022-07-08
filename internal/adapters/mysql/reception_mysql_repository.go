package mysql

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/process"
	"github.com/alejandroik/trazavino-api/internal/domain/reception"
	"gorm.io/gorm"
)

type ReceptionModel struct {
	gorm.Model
	ProcessID uint
	Process   ProcessModel
	//TruckID   uint
	//Truck     vy.Truck
	//vineyard  vy.Vineyard
	//grapeType vy.GrapeType
	Weight int64
	Sugar  int64
}

func (ReceptionModel) TableName() string {
	return "reception"
}

type ReceptionMysqlRepository struct {
	db *gorm.DB
}

func NewReceptionMysqlRepository(db *gorm.DB) *ReceptionMysqlRepository {
	if db == nil {
		panic("missing db")
	}

	return &ReceptionMysqlRepository{db: db}
}

func (r ReceptionMysqlRepository) AddReception(ctx context.Context, rc *reception.Reception) error {
	t := time.Now()
	pm := ProcessModel{
		StartDate: &t,
		Ptype:     process.Reception.String(),
	}
	err := addProcess(r.db, &pm)
	if err != nil {
		return err
	}

	rm := r.marshallReception(rc, pm)
	result := r.db.Create(&rm)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r ReceptionMysqlRepository) GetReception(ctx context.Context, receptionId int64) (*reception.Reception, error) {
	rm := ReceptionModel{}
	result := r.db.First(&rm, receptionId)
	if result.Error != nil {
		return nil, result.Error
	}
	rec, err := r.unmarshallReception(rm)
	if err != nil {
		return nil, err
	}

	return rec, nil
}

func (r ReceptionMysqlRepository) GetAllReceptions() ([]*reception.Reception, error) {
	return nil, nil
}

func (r ReceptionMysqlRepository) UpdateReception(ctx context.Context, receptionId int64, updateFn func(ctx context.Context, rc *reception.Reception) (*reception.Reception, error)) {
}

func (r ReceptionMysqlRepository) marshallReception(rc *reception.Reception, process ProcessModel) ReceptionModel {
	rm := ReceptionModel{
		ProcessID: process.ID,
		//Truck:     rc.Truck(),
		Weight: rc.Weight(),
		Sugar:  rc.Sugar(),
	}

	return rm
}

func (r ReceptionMysqlRepository) unmarshallReception(rm ReceptionModel) (*reception.Reception, error) {
	return reception.UnmarshallReceptionFromDatabase(rm.Weight, rm.Sugar)
}
