package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ashsajal1/go-tweet/auth"
	"github.com/ashsajal1/go-tweet/controller"
	"github.com/ashsajal1/go-tweet/initializers"
	"github.com/ashsajal1/go-tweet/model"
)

func init() {
	initializers.LoadEnv()
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

	auth.SetupAuthRouter(api)     
	controller.SetupUserRoute(api) 
	controller.SetupTweetRoute(api) 
	controller.SetupLikeRoute(api)  

	r.Run(":8080")
}
