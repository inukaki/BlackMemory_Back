package repository

import (
	"go_rest_api/model"

	"gorm.io/gorm"
)

type IWorkRepository interface {
	CreateWork(work *model.Work) error
}

type workRepository struct {
	db *gorm.DB
}

func NewWorkRepository(db *gorm.DB) IWorkRepository {
	return &workRepository{db}
}

func (wr *workRepository) CreateWork(work *model.Work) error {
	if err := wr.db.Create(work).Error; err != nil {
		return err
	}
	return nil
}
