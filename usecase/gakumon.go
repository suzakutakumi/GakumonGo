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

type GakumonUsecaseImpl struct {
	r GakumonRepository
}

func NewGakumonUsecase(r GakumonRepository) *GakumonUsecaseImpl {
	return &GakumonUsecaseImpl{r: r}
}

func (u GakumonUsecaseImpl) GetGakumonList() ([]model.GakumonId, error) {
	return u.r.FetchAllGakumonId()
}
