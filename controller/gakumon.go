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
