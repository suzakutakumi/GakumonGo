package usecase

import (
	"fmt"
	"gakumon_go/model"
	"mime/multipart"
)

type GakumonUsecase interface {
	GetGakumonList() ([]model.GakumonID, error)
	RegisterNewGakumon(model.Gakumon, multipart.File, multipart.FileHeader) (model.Gakumon, error)
}

type GakumonRepository interface {
	FetchAllGakumonID() ([]model.GakumonID, error)
	InsertGakumonInfo(model.Gakumon) (model.Gakumon, error)
	UploadGakumonImage(multipart.File, multipart.FileHeader) (string, error)
	DeleteGakumonImage(string) (string, error)
}

type gakumonUsecaseImpl struct {
	r GakumonRepository
}

func NewGakumonUsecase(r GakumonRepository) *gakumonUsecaseImpl {
	return &gakumonUsecaseImpl{r: r}
}

func (u gakumonUsecaseImpl) GetGakumonList() ([]model.GakumonID, error) {
	return u.r.FetchAllGakumonID()
}

func (u gakumonUsecaseImpl) RegisterNewGakumon(g model.Gakumon, img multipart.File, img_info multipart.FileHeader) (model.Gakumon, error) {
	path, err := u.r.UploadGakumonImage(img, img_info)
	if err != nil {
		return model.Gakumon{}, err
	}

	g.ImageURL = &path

	gakumon, err := u.r.InsertGakumonInfo(g)
	if err != nil {
		if _, derr := u.r.DeleteGakumonImage(img_info.Filename); derr != nil {
			return model.Gakumon{}, fmt.Errorf("%s&%s", err.Error(), derr.Error())
		}
		return model.Gakumon{}, err
	}

	return gakumon, nil
}
