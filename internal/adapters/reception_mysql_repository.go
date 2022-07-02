package adapters

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/reception"
	"gorm.io/gorm"
)

type ReceptionModel struct {
	ProcessModel
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
	return nil
}

func (r ReceptionMysqlRepository) GetReception(ctx context.Context, receptionId int64) (*reception.Reception, error) {
	return nil, nil
}

//func (r ReceptionMysqlRepository) GetAllReceptions() ([]*reception.Reception, error) {
//	var receptions []*ReceptionModel
//	if err := r.db.Find(&receptions).Error; err != nil {
//		return nil, err
//	}
//
//	return r.unmarshallReceptions(receptions)
//}

func (r ReceptionMysqlRepository) UpdateReception(ctx context.Context, receptionId int64, updateFn func(ctx context.Context, rc *reception.Reception) (*reception.Reception, error)) {
}

func (r ReceptionMysqlRepository) unmarshallReception() (*reception.Reception, error) {
	return reception.UnmarshallReceptionFromDatabase()
}
