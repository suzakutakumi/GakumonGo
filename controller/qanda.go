package controller

import (
	"gakumon_go/model"
	"gakumon_go/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionController interface {
	GetQuestion()
}

type AnswerController interface {
	CheckCollect()
}

type QuestionControllerImpl struct {
	u usecase.QuestionUsecase
}

type AnswerControllerImpl struct {
	u usecase.AnswerUsecase
}

func NewQuestionControllerImpl(u usecase.QuestionUsecase) *QuestionControllerImpl {
	return &QuestionControllerImpl{u: u}
}

func NewAnswerControllerImpl(u usecase.AnswerUsecase) *AnswerControllerImpl {
	return &AnswerControllerImpl{u: u}
}

func (c QuestionControllerImpl) GetQuestion(ctx *gin.Context) {
	id := ctx.Param("gakumon_id")
	l, err := c.u.GetQuestion(id)
	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, l)
}

func (c AnswerControllerImpl) CheckCollect(ctx *gin.Context) {
	id := ctx.Param("gakumon_id")
	var j model.Answer
	err := ctx.ShouldBindJSON(&j)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	l, err := c.u.CheckCollect(id, j)
	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, l)
}
