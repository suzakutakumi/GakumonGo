package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetGakumonIdList(t *testing.T) {

	if err := godotenv.Load("../.env"); err != nil {
		t.Error("環境変数が読み取れませんでした")
		return
	}
	keypath := "../" + os.Getenv("ACCOUNT_KEY_JSON_PATH")

	ctx := context.Background()

	app, err := InitFirebase(ctx, keypath)
	if err != nil {
		log.Fatal(err.Error())
	}

	storeCli, err := InitFirestore(ctx, app)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer storeCli.Close()

	r := GakumonRepositoryImpl{
		Firestore: storeCli,
		Ctx:       ctx,
	}
	studentIdList, err := r.FetchAllGakumonID()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(studentIdList)
}
