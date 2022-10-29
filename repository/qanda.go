package repository

import (
	"context"
	"gakumon_go/model"

	"cloud.google.com/go/firestore"
)

type QuestionRepositoryImpl struct {
	Client *firestore.Client
	Ctx    context.Context
}

type AnswerRepositoryImpl struct {
	Client *firestore.Client
	Ctx    context.Context
}

func (r QuestionRepositoryImpl) FindQuestion(id string) (model.Question, error) {
	var question model.Question
	dsnap, err := r.Client.Collection("QANDA").Doc(id).Get(r.Ctx)
	if err != nil {
		return model.Question{}, err
	}
	err = dsnap.DataTo(&question)
	if err != nil {
		return model.Question{}, err
	}
	return question, nil
}

func (r AnswerRepositoryImpl) FindAnswer(id string) (model.Answer, error) {
	var answer model.Answer
	dsnap, err := r.Client.Collection("QANDA").Doc(id).Get(r.Ctx)
	if err != nil {
		return model.Answer{}, err
	}
	err = dsnap.DataTo(&answer)
	if err != nil {
		return model.Answer{}, err
	}
	return answer, nil
}

func (r AnswerRepositoryImpl) FindGakumon(id string) (model.Gakumon, error) {
	var answer model.Gakumon
	dsnap, err := r.Client.Collection("GAKUMON").Doc(id).Get(r.Ctx)
	if err != nil {
		return model.Gakumon{}, err
	}
	err = dsnap.DataTo(&answer)
	if err != nil {
		return model.Gakumon{}, err
	}
	return answer, nil
}
