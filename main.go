package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ashsajal1/go-tweet/controller"
	"github.com/ashsajal1/go-tweet/model"
	"github.com/ashsajal1/go-tweet/auth"
)

func init() {
	model.InitializeDB()
}

func main() {
	r := gin.Default()

	api := r.Group("/api") // Create a new route group for the API

	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	auth.SetupAuthRouter(api)     // Change to use the api group
	controller.SetupUserRoute(api) // Change to use the api group
	controller.SetupTweetRoute(api) // Change to use the api group
	controller.SetupLikeRoute(api)  // Change to use the api group

	r.Run(":8080")
}
