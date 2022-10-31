package usecase

import (
	"fmt"
	"gakumon_go/model"
	"mime/multipart"
)

type GakumonUsecase interface {
	GetGakumonList() ([]model.GakumonID, error)
	RegisterNewGakumon(model.Gakumon, model.QandA, *multipart.File, *string) error
}

type GakumonRepository interface {
	FetchAllGakumonID() ([]model.GakumonID, error)
	InsertGakumonInfo(model.Gakumon) error
	InsertQandA(model.QandA) error
	DeleteQandA(string) error
	UploadGakumonImage(multipart.File, string) (string, error)
	DeleteGakumonImage(string) error
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

func (u gakumonUsecaseImpl) RegisterNewGakumon(g model.Gakumon, qanda model.QandA, img *multipart.File, img_name *string) error {
	if err := u.r.InsertQandA(qanda); err != nil {
		return err
	}

	if img == nil || img_name == nil {
		if err := u.r.InsertGakumonInfo(g); err != nil {
			if dqerr := u.r.DeleteQandA(qanda.GakumonID); err != nil {
				return fmt.Errorf("%s&%s", err.Error(), dqerr.Error())
			}
			return err
		}
	} else {
		path, err := u.r.UploadGakumonImage(*img, *img_name)
		if err != nil {
			return err
		}

		g.ImageURL = &path

		if err := u.r.InsertGakumonInfo(g); err != nil {
			errMsg := err.Error()
			if dgerr := u.r.DeleteGakumonImage(*img_name); dgerr != nil {
				errMsg += "&" + dgerr.Error()
			}
			if dqerr := u.r.DeleteQandA(qanda.GakumonID); err != nil {
				errMsg += "&" + dqerr.Error()
			}
			return fmt.Errorf(errMsg)
		}
	}

	return nil
}
