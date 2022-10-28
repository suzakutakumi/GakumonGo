package usecase

import (
	"gakumon_go/model"
)

type GakumonUsecase interface {
	GetGakumonList() []model.GakumonId
}

type GakumonRepository interface {
	FetchAllGakumonId() []model.GakumonId
}

type GakumonUsecaseImpl struct {
	r GakumonRepository
}

func NewGakumonUsecase(r GakumonRepository) *GakumonUsecaseImpl {
	return &GakumonUsecaseImpl{r: r}
}

func (u GakumonUsecaseImpl) GetGakumonList() []model.GakumonId {
	return u.r.FetchAllGakumonId()
}
