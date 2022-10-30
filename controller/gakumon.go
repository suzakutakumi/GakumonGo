package controller

import (
	"encoding/json"
	"gakumon_go/model"
	"gakumon_go/usecase"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GakumonController interface {
	GetGakumonList()
	RegisterNewGakumon()
}

type gakumonControllerImpl struct {
	u usecase.GakumonUsecase
}

func NewGakumonController(u usecase.GakumonUsecase) *gakumonControllerImpl {
	return &gakumonControllerImpl{u: u}
}

func (c gakumonControllerImpl) GetGakumonList(ctx *gin.Context) {
	l, err := c.u.GetGakumonList()
	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, l)
}

func (c gakumonControllerImpl) RegisterNewGakumon(ctx *gin.Context) {
	var gakumon model.Gakumon
	gakumonJson := ctx.Request.FormValue("gakumon")
	if err := json.Unmarshal([]byte(gakumonJson), &gakumon); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var qanda model.QandA
	qandaJson := ctx.Request.FormValue("qanda")
	if err := json.Unmarshal([]byte(qandaJson), &qanda); err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	var imgPtr *multipart.File = nil
	var imgNamePtr *string = nil

	img, imgInfo, err := ctx.Request.FormFile("image")
	if err == nil {
		imgName := gakumon.GakumonID

		imgType := imgInfo.Header.Get("Content-Type")
		if imgType == "image/jpeg" {
			imgName += ".jpg"
		} else if imgType == "image/png" {
			imgName += ".png"
		} else {
			log.Println(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		imgPtr = &img
		imgNamePtr = &imgName
	}
	if err := c.u.RegisterNewGakumon(gakumon, qanda, imgPtr, imgNamePtr); err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusCreated)
}
