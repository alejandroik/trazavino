package gorm

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"gorm.io/gorm"
)

type TruckModel struct {
	BaseEntityModel
}

func (TruckModel) TableName() string {
	return "truck"
}

type TruckMysqlRepository struct {
	db *gorm.DB
}

func NewTruckMysqlRepository(db *gorm.DB) *TruckMysqlRepository {
	if db == nil {
		panic("missing db")
	}

	return &TruckMysqlRepository{db: db}
}

func (r TruckMysqlRepository) AddTruck(ctx context.Context, truck *entity.Truck) (*entity.Truck, error) {
	tm := r.marshalTruck(truck)
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&tm).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	savedTruck, err := r.unmarshalTruck(tm)
	if err != nil {
		return nil, err
	}

	return savedTruck, nil
}

func (r TruckMysqlRepository) GetTruck(ctx context.Context, truckId int) (*entity.Truck, error) {
	return nil, nil
}

func (r TruckMysqlRepository) marshalTruck(truck *entity.Truck) TruckModel {
	return TruckModel{
		BaseEntityModel{
			Name: truck.Name(),
		},
	}
}

func (r TruckMysqlRepository) unmarshalTruck(tm TruckModel) (*entity.Truck, error) {
	return entity.UnmarshalTruckFromDatabase(int(tm.ID), tm.Name)
}
