package repository

import (
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

	client, ctx, err := InitFirebase(keypath)
	if err != nil {
		t.Error("firebaseの初期化ができませんでした")
		return
	}
	defer client.Close()

	r := GakumonRepositoryImpl{
		Client: client,
		Ctx:    ctx,
	}
	studentIdList, err := r.FetchAllGakumonId()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(studentIdList)
}
