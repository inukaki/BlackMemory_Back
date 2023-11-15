package usecase

import (
	"go_rest_api/model"
	"go_rest_api/repository"
)

type IWorkUsecase interface {
	CreateWork(work model.Work) (model.WorkResponse, error)
	UpdateWork(work model.Work, userId uint, workDate string) (model.WorkResponse, error)
	GetWorkByDate(userId uint, workDate string) (model.WorkResponse, error)
}

type workUsecase struct {
	wr repository.IWorkRepository
}

func NewWorkUseCase(wr repository.IWorkRepository) IWorkUsecase {
	return &workUsecase{wr}
}

func (wu *workUsecase) CreateWork(work model.Work) (model.WorkResponse, error) {
	if err := wu.wr.CreateWork(&work); err != nil {
		return model.WorkResponse{}, err
	}
	resWork := model.WorkResponse{
		ID:      work.ID,
		Date:    work.Date,
		StartAt: work.StartAt,
		EndAt:   work.EndAt,
		Hours:   work.Hours,
		Content: work.Content,
	}
	return resWork, nil
}

func (wu *workUsecase) UpdateWork(work model.Work, userId uint, workDate string) (model.WorkResponse, error) {
	if err := wu.wr.UpdateWork(&work, userId, workDate); err != nil {
		return model.WorkResponse{}, err
	}
	resWork := model.WorkResponse{
		ID:      work.ID,
		Date:    work.Date,
		StartAt: work.StartAt,
		EndAt:   work.EndAt,
		Hours:   work.Hours,
		Content: work.Content,
	}
	return resWork, nil
}

func (wu *workUsecase) GetWorkByDate(workId uint, workDate string) (model.WorkResponse, error) {
	work := model.Work{}
	if err := wu.wr.GetWorkByDate(&work, workId, workDate); err != nil {
		return model.WorkResponse{}, err
	}
	resWork := model.WorkResponse{
		ID:      work.ID,
		Date:    work.Date,
		StartAt: work.StartAt,
		EndAt:   work.EndAt,
		Hours:   work.Hours,
		Content: work.Content,
	}
	return resWork, nil
}
