package gorm

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/alejandroik/trazavino-api/internal/domain/entity/enum/process_type"
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
	rm := ReceptionModel{
		Process: ProcessModel{
			StartDate: time.Now(),
			PType:     process_type.Reception.String(),
		},
		TruckID: uint(rc.TruckID()),
		Weight:  int(rc.Weight()),
		Sugar:   int(rc.Sugar()),
	}

	return rm
}

func (r ReceptionMysqlRepository) unmarshallReception(rm ReceptionModel) (*entity.Reception, error) {
	return entity.UnmarshalReceptionFromDatabase(int64(rm.Process.ID), int64(rm.TruckID), 0, 0, int32(rm.Weight), int32(rm.Sugar))
}
