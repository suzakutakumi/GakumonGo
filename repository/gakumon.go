package repository

import (
	"context"
	"errors"
	"gakumon_go/model"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func InitFirebase(path string) (*firestore.Client, context.Context, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, nil, errors.New("firebaseの初期化に失敗しました")
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, nil, errors.New("firestoreの初期化に失敗しました")
	}

	return client, ctx, nil
}

type GakumonRepositoryImpl struct {
	Client *firestore.Client
	Ctx    context.Context
}

func (r GakumonRepositoryImpl) FetchAllGakumonId() ([]model.GakumonId, error) {
	var gakumon_id_list []model.GakumonId
	iter := r.Client.Collection("GAKUMON").Documents(r.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var id model.GakumonId
		if err := doc.DataTo(&id); err != nil {
			return nil, err
		}
		gakumon_id_list = append(gakumon_id_list, id)

	}
	return gakumon_id_list, nil

}
