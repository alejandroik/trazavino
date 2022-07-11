package gorm

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReceptionModel struct {
	gorm.Model
	ProcessID uint
	Process   ProcessModel
	TruckID   uint
	Truck     TruckModel
	//VineyardID  uint
	//Vineyard    VineyardModel
	//GrapeTypeID uint
	//GrapeType   GrapeTypeModel
	Weight int
	Sugar  int
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

func (r ReceptionMysqlRepository) AddReception(ctx context.Context, rc *entity.Reception) (*entity.Reception, error) {
	rm := r.marshallReception(rc)
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&rm).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	rec, err := r.unmarshallReception(rm)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func (r ReceptionMysqlRepository) GetReception(ctx context.Context, receptionId int) (*entity.Reception, error) {
	rm := ReceptionModel{}
	if err := r.db.Preload(clause.Associations).First(&rm, receptionId).Error; err != nil {
		return nil, err
	}
	rec, err := r.unmarshallReception(rm)
	if err != nil {
		return nil, err
	}

	return rec, nil
}

func (r ReceptionMysqlRepository) GetAllReceptions() ([]*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionMysqlRepository) UpdateReception(ctx context.Context, receptionId int, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		//if err := tx.Model()

		//db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
		return nil
	})
}

func (r ReceptionMysqlRepository) marshallReception(rc *entity.Reception) ReceptionModel {
	t := time.Now()
	rm := ReceptionModel{
		Process: ProcessModel{
			StartDate: &t,
			Ptype:     entity.TypeReception.String(),
		},
		TruckID: uint(rc.Truck().ID()),
		Weight:  rc.Weight(),
		Sugar:   rc.Sugar(),
	}

	return rm
}

func (r ReceptionMysqlRepository) unmarshallReception(rm ReceptionModel) (*entity.Reception, error) {
	process, err := entity.UnmarshalProcessFromDatabase(int(rm.Process.ID), rm.Process.StartDate, rm.Process.EndDate, rm.Process.Ptype, rm.Process.Hash, rm.Process.Transaction, rm.Process.Temperature)
	if err != nil {
		return nil, err
	}
	truck, err := entity.UnmarshalTruckFromDatabase(int(rm.Truck.ID), rm.Truck.Name)
	if err != nil {
		return nil, err
	}
	return entity.UnmarshalReceptionFromDatabase(process, truck, rm.Weight, rm.Sugar)
}
