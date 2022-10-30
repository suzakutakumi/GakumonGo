package main

import (
	"context"
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
		log.Fatalf("環境変数がない:%s", err.Error())
	}
	keypath := os.Getenv("ACCOUNT_KEY_JSON_PATH")

	ctx := context.Background()

	app, err := repository.InitFirebase(ctx, keypath)
	if err != nil {
		log.Fatal(err.Error())
	}

	storeCli, err := repository.InitFirestore(ctx, app)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer storeCli.Close()

	storageCli, err := repository.InitCloudStorage(ctx, app)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := repository.GakumonRepositoryImpl{Firestore: storeCli, Storage: storageCli, Ctx: ctx}
	u := usecase.NewGakumonUsecase(r)
	gakumonctrl := controller.NewGakumonController(u)

	questionctrl := controller.NewQuestionControllerImpl(
		usecase.NewQuestionUsecase(
			repository.QuestionRepositoryImpl{Client: storeCli, Ctx: ctx},
		),
	)

	answerctrl := controller.NewAnswerControllerImpl(
		usecase.NewAnswerUsecase(
			repository.AnswerRepositoryImpl{Client: storeCli, Ctx: ctx},
		),
	)
	router := gin.Default()
	router.GET("/gakumon_list", gakumonctrl.GetGakumonList)
	router.POST("/gakumon", gakumonctrl.RegisterNewGakumon)
	router.GET("/gakumon/:gakumon_id/question", questionctrl.GetQuestion)
	router.POST("/gakumon/:gakumon_id/answer", answerctrl.CheckCollect)

	router.Run(":8080")
}
