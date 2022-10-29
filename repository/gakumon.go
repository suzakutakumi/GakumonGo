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
	client *firestore.Client
	ctx    context.Context
}

func (r GakumonRepositoryImpl) FetchAllGakumonId() []model.GakumonId {
	var gakumon_id_list []model.GakumonId
	iter := r.client.Collection("GAKUMON").Documents(r.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []model.GakumonId{{GakumonId: err.Error()}}
		}
		var id model.GakumonId
		if err := doc.DataTo(&id); err != nil {
			return []model.GakumonId{{GakumonId: err.Error()}}
		}
		gakumon_id_list = append(gakumon_id_list, id)

	}
	return gakumon_id_list

}
