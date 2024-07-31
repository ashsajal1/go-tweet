package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ashsajal1/go-tweet/controller"
	"github.com/ashsajal1/go-tweet/model"
)

func init() {
	model.InitializeDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	controller.SetupUserRoute(r)
	controller.SetupTweetRoute(r)
	controller.SetupLikeRoute(r)

	r.Run(":8080")
}
