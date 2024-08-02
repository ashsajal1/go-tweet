package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupLikeRoute(r *gin.RouterGroup) {
	like := r.Group("/like")
	{
		like.POST("", CreateLike)
		like.DELETE("/:id", DeleteLike)
		like.GET("/:id", GetLike)
	}
}

func SetupTweetRoute(r *gin.RouterGroup) {
	tweet := r.Group("/tweet")
	{
		tweet.POST("", CreateTweet)
		tweet.GET("/:id", GetTweet)
		tweet.GET("", GetTweets)
	}
}

func SetupUserRoute(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("", CreateUser)
		user.GET("/:id", GetUser)
		user.GET("/:id/tweets", GetTweetsByUserID)
		user.GET("", GetUsers)
	}
}
