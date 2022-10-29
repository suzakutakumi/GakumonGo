package usecase

import (
	"gakumon_go/model"
)

type GakumonUsecase interface {
	GetGakumonList() ([]model.GakumonID, error)
}

type GakumonRepository interface {
	FetchAllGakumonId() ([]model.GakumonID, error)
}

type gakumonUsecaseImpl struct {
	r GakumonRepository
}

func NewGakumonUsecase(r GakumonRepository) *gakumonUsecaseImpl {
	return &gakumonUsecaseImpl{r: r}
}

func (u gakumonUsecaseImpl) GetGakumonList() ([]model.GakumonID, error) {
	return u.r.FetchAllGakumonId()
}
