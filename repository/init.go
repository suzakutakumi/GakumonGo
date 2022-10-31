package repository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

func InitFirebase(ctx context.Context, path string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.New("firebaseの初期化に失敗しました")
	}
	return app, nil
}

func InitFirestore(ctx context.Context, app *firebase.App) (*firestore.Client, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.New("firestoreの初期化に失敗しました")
	}

	return client, nil
}

func InitCloudStorage(ctx context.Context, app *firebase.App) (*storage.Client, error) {
	client, err := app.Storage(ctx)
	if err != nil {
		return nil, errors.New("cloud storageの初期化に失敗しました")
	}

	return client, nil
}
