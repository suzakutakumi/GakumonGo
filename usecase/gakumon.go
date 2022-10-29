package usecase

import (
	"gakumon_go/model"
)

type GakumonUsecase interface {
	GetGakumonList() ([]model.GakumonId, error)
}

type GakumonRepository interface {
	FetchAllGakumonId() ([]model.GakumonId, error)
}

type gakumonUsecaseImpl struct {
	r GakumonRepository
}

func NewGakumonUsecase(r GakumonRepository) *gakumonUsecaseImpl {
	return &gakumonUsecaseImpl{r: r}
}

func (u gakumonUsecaseImpl) GetGakumonList() ([]model.GakumonId, error) {
	return u.r.FetchAllGakumonId()
}
