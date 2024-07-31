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

	user := r.Group("/user")
	{
		user.POST("", controller.CreateUser)
		user.GET("/:id", controller.GetUser)
		user.GET("/:id/tweets", controller.GetTweetsByUserID)
		user.GET("", controller.GetUsers)
	}

	tweet := r.Group("/tweet")
	{
		tweet.POST("", controller.CreateTweet)
		tweet.GET("/:id", controller.GetTweet)
		tweet.GET("", controller.GetTweets)
	}

	r.Run(":8080")
}
