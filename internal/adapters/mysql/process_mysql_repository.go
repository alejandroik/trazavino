package mysql

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/process"
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

func (r ProcessMysqlRepository) GetProcess(ctx context.Context, processId uint) (*process.Process, error) {
	pm := ProcessModel{}
	result := r.db.First(&pm, processId)
	if result.Error != nil {
		return nil, result.Error
	}
	pr, err := unmarshallProcess(pm)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

func (r ProcessMysqlRepository) GetAllProcesses() ([]*process.Process, error) {
	return nil, nil
}

func addProcess(db *gorm.DB, pm *ProcessModel) error {
	result := db.Create(pm)
	if result.Error != nil {
		return result.Error
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

func unmarshallProcess(pm ProcessModel) (*process.Process, error) {
	return process.UnmarshallProcessFromDatabase(pm.ID, *pm.StartDate, *pm.EndDate, pm.Ptype, pm.Hash, pm.Transaction, pm.Temperature)
}
