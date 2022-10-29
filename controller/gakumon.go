package controller

import (
	"gakumon_go/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GakumonController interface {
	GetGakumonList()
}

type GakumonControllerImp struct {
	u usecase.GakumonUsecase
}

func NewGakumonController(u usecase.GakumonUsecase) *GakumonControllerImp {
	return &GakumonControllerImp{u: u}
}

func (c GakumonControllerImp) GetGakumonList(ctx *gin.Context) {
	l, err := c.u.GetGakumonList()
	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, l)
}
