package gorm

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"gorm.io/gorm"
)

type MacerationModel struct {
	gorm.Model
	ProcessID uint
}

type MacerationMysqlRepository struct {
	db *gorm.DB
}

func NewMacerationMysqlRepository(db *gorm.DB) *MacerationMysqlRepository {
	return &MacerationMysqlRepository{db: db}
}

func (r MacerationMysqlRepository) AddMaceration(ctx context.Context, m *entity.Maceration) error {
	return nil
}

func (r MacerationMysqlRepository) GetMaceration(ctx context.Context, macerationId int64) (*entity.Maceration, error) {
	return nil, nil
}

func (r MacerationMysqlRepository) UpdateMaceration(ctx context.Context, macerationId int64, updateFn func(ctx context.Context, m *entity.Maceration) (*entity.Maceration, error)) {
}
