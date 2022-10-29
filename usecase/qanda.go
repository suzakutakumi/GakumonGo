package usecase

import (
	"gakumon_go/model"
)

type QuestionUsecase interface {
	GetQuestion(string) (model.Question, error)
}

type QuestionRepository interface {
	FindQuestion(string) (model.Question, error)
}

type AnswerUsecase interface {
	CheckCollect(string, model.Answer) (model.GakumonByAnswer, error)
}

type AnswerRepository interface {
	FindAnswer(string) (model.Answer, error)
	FindGakumon(string) (model.Gakumon, error)
}

type QuestionUsecaseImpl struct {
	r QuestionRepository
}

type AnswerUsecaseImpl struct {
	r AnswerRepository
}

func NewQuestionUsecase(r QuestionRepository) *QuestionUsecaseImpl {
	return &QuestionUsecaseImpl{r: r}
}

func NewAnswerUsecase(r AnswerRepository) *AnswerUsecaseImpl {
	return &AnswerUsecaseImpl{r: r}
}

func (u QuestionUsecaseImpl) GetQuestion(id string) (model.Question, error) {
	return u.r.FindQuestion(id)
}

func (u AnswerUsecaseImpl) CheckCollect(id string, ans model.Answer) (model.GakumonByAnswer, error) {
	correctAnswer, err := u.r.FindAnswer(id)
	if err != nil {
		return model.GakumonByAnswer{}, nil
	}
	if correctAnswer == ans {
		gakumon, err := u.r.FindGakumon(id)
		if err != nil {
			return model.GakumonByAnswer{}, nil
		}
		return model.GakumonByAnswer{
			IsCollect: true,
			User:      &gakumon,
		}, nil
	}
	return model.GakumonByAnswer{
		IsCollect: false,
		User:      nil,
	}, nil
}
