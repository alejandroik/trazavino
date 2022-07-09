package mysql

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/process"
	"github.com/alejandroik/trazavino-api/internal/domain/reception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	return r.db.Transaction(func(tx *gorm.DB) error {
		t := time.Now()
		pm := ProcessModel{StartDate: &t, Ptype: process.Reception.String()}
		if err := addProcess(tx, &pm); err != nil {
			return err
		}

		rm := r.marshallReception(rc, pm)
		if err := r.db.Create(&rm).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r ReceptionMysqlRepository) GetReception(ctx context.Context, receptionId int64) (*reception.Reception, error) {
	rm := ReceptionModel{}
	result := r.db.Preload(clause.Associations).First(&rm, receptionId)
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

func (r ReceptionMysqlRepository) UpdateReception(ctx context.Context, receptionId int64, updateFn func(ctx context.Context, rc *reception.Reception) (*reception.Reception, error)) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		//if err := tx.Model()

		//db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
		return nil
	})
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
	p, err := unmarshallProcess(rm.Process)
	if err != nil {
		return nil, err
	}
	return reception.UnmarshallReceptionFromDatabase(p, rm.Weight, rm.Sugar)
}
