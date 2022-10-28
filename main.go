package main

import (
	"gakumon_go/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/gakumon_list", controller.GetGakumonList)
	router.Run(":8080")
}
