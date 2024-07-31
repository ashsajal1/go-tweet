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
	

	r.POST("/user", controller.CreateUser)
	r.GET("/user/:id", controller.GetUser)
	r.GET("/user/:id/tweets", controller.GetTweetsByUserID)
	r.GET("/user", controller.GetUsers)

	r.POST("/tweet", controller.CreateTweet)
	r.GET("/tweet/:id", controller.GetTweet)
	r.GET("/tweet", controller.GetTweets)

	r.Run(":8080")
}	