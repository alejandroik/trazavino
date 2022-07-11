package gorm

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"gorm.io/gorm"
)

type ProcessModel struct {
	gorm.Model
	StartDate   *time.Time
	EndDate     *time.Time
	Hash        string
	Transaction string
	Ptype       string
	Temperature int
}

func (ProcessModel) TableName() string {
	return "process"
}

type ProcessMysqlRepository struct {
	db *gorm.DB
}

func NewProcessMysqlRepository(db *gorm.DB) *ProcessMysqlRepository {
	if db == nil {
		panic("missing db")
	}

	return &ProcessMysqlRepository{db: db}
}

func (r ProcessMysqlRepository) GetProcess(ctx context.Context, processId uint) (*entity.Process, error) {
	pm := ProcessModel{}
	result := r.db.First(&pm, processId)
	if result.Error != nil {
		return nil, result.Error
	}
	pr, err := unmarshalProcess(pm)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

func (r ProcessMysqlRepository) GetAllProcesses() ([]*entity.Process, error) {
	return nil, nil
}

func addProcess(db *gorm.DB, pm *ProcessModel) error {
	err := db.Create(pm).Error
	if err != nil {
		return err
	}

	return nil
}

//func marshallProcess(pr process.Process) ProcessModel {
//	pm := ProcessModel{
//		StartDate:   pr.StartDate(),
//		EndDate:     pr.EndDate(),
//		Hash:        pr.Hash(),
//		Transaction: pr.Transaction(),
//		Ptype:       pr.Ptype(),
//		Temperature: pr.Temperature(),
//	}
//
//	return pm
//}

func unmarshalProcess(pm ProcessModel) (*entity.Process, error) {
	return entity.UnmarshalProcessFromDatabase(int(pm.ID), pm.StartDate, pm.EndDate, pm.Ptype, pm.Hash, pm.Transaction, pm.Temperature)
}
