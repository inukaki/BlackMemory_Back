package repository

import (
	"fmt"
	"go_rest_api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IWorkRepository interface {
	CreateWork(work *model.Work) error
	UpdateWork(work *model.Work, userId uint, workDate string) error
	GetWorkByDate(work *model.Work, userId uint, workDate string) error
	GetAllWorks(works *[]model.Work, userId uint) error
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

func (wr *workRepository) UpdateWork(work *model.Work, userId uint, workDate string) error {
	result := wr.db.Model(work).Clauses(clause.Returning{}).Where("user_id = ? AND date = ?", userId, workDate).Updates(map[string]interface{}{"start_at": work.StartAt, "end_at": work.EndAt, "hours": work.Hours, "content": work.Content})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (wr *workRepository) GetWorkByDate(work *model.Work, userId uint, workDate string) error {
	if err := wr.db.Where("user_id = ? AND date = ?", userId, workDate).First(work).Error; err != nil {
		return err
	}
	return nil
}

func (wr *workRepository) GetAllWorks(works *[]model.Work, userId uint) error {
	if err := wr.db.Where("user_id = ?", userId).Order("date").Find(works).Error; err != nil {
		return err
	}
	return nil
}
