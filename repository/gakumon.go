package repository

import (
	"context"
	"fmt"
	"gakumon_go/model"
	"io"
	"mime/multipart"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type GakumonRepositoryImpl struct {
	Firestore *firestore.Client
	Storage   *storage.Client
	Ctx       context.Context
}

func (r GakumonRepositoryImpl) FetchAllGakumonID() ([]model.GakumonID, error) {
	var gakumon_id_list []model.GakumonID
	iter := r.Firestore.Collection("GAKUMON").Documents(r.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var id model.GakumonID
		if err := doc.DataTo(&id); err != nil {
			return nil, err
		}
		gakumon_id_list = append(gakumon_id_list, id)

	}

	return gakumon_id_list, nil
}

func (r GakumonRepositoryImpl) InsertGakumonInfo(g model.Gakumon) error {
	_, err := r.Firestore.Collection("GAKUMON").Doc(g.GakumonID).Set(r.Ctx, g)
	if err != nil {
		return err
	}
	return nil
}

func (r GakumonRepositoryImpl) UploadGakumonImage(file multipart.File, fileName string) (string, error) {
	bucket, err := r.Storage.Bucket("gakumongo-ae7fb.appspot.com")
	if err != nil {
		return "", err
	}

	o := bucket.Object(fileName)
	w := o.NewWriter(r.Ctx)
	if _, err = io.Copy(w, file); err != nil {
		return "", err
	}
	if err = w.Close(); err != nil {
		return "", err
	}

	return fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/gakumongo-ae7fb.appspot.com/o/%s?alt=media", fileName), nil
}

func (r GakumonRepositoryImpl) DeleteGakumonImage(fileName string) error {
	bucket, err := r.Storage.Bucket("gakumongo-ae7fb.appspot.com")
	if err != nil {
		return err
	}

	err = bucket.Object(fileName).Delete(r.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r GakumonRepositoryImpl) InsertQandA(qanda model.QandA) error {
	_, err := r.Firestore.Collection("QANDA").Doc(qanda.GakumonID).Set(r.Ctx, qanda)
	if err != nil {
		return err
	}
	return nil
}

func (r GakumonRepositoryImpl) DeleteQandA(id string) error {
	_, err := r.Firestore.Collection("QANDA").Doc(id).Delete(r.Ctx)
	if err != nil {
		return err
	}
	return nil
}
