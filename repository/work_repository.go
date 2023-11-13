package repository

import (
	"fmt"
	"go_rest_api/model"

	"gorm.io/gorm"
)

type IWorkRepository interface {
	CreateWork(work *model.Work) error
	UpdateWork(work *model.Work, userId uint, workId uint) error
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

func (wr *workRepository) UpdateWork(work *model.Work, userId uint, workId uint) error {
	result := wr.db.Model(work).Where("id = ? AND user_id = ?", workId, userId).Updates(map[string]interface{}{"start_at": work.StartAt, "end_at": work.EndAt, "hours": work.Hours, "content": work.Content})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
