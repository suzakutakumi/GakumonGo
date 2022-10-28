package controller

import (
	"gakumon_go/repository"
	"gakumon_go/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGakumonList(c *gin.Context) {
	u := usecase.NewGakumonUsecase(repository.GakumonRepositoryImpl{})
	l := u.GetGakumonList()
	c.JSON(http.StatusOK, l)
}
