package main

import (
	"gakumon_go/controller"
	"gakumon_go/repository"
	"gakumon_go/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("環境変数がない")
	}
	keypath := os.Getenv("ACCOUNT_KEY_JSON_PATH")

	client, ctx, err := repository.InitFirebase(keypath)
	if err != nil {
		log.Fatal("firebaseが初期化されなかった")
	}
	defer client.Close()

	r := repository.GakumonRepositoryImpl{Client: client, Ctx: ctx}
	u := usecase.NewGakumonUsecase(r)
	c := controller.NewGakumonController(u)

	router := gin.Default()
	router.GET("/gakumon_list", c.GetGakumonList)
	router.Run(":8080")
}
