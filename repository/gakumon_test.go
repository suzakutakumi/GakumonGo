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
		client: client,
		ctx:    ctx,
	}
	studentIdList := r.FetchAllGakumonId()
	t.Log(studentIdList)
}
