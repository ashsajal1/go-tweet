package controller

import (
	"errors"
	"strconv"

	"github.com/ashsajal1/go-tweet/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTweet(c *gin.Context) {
	var tweet model.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	db.Create(&tweet)
	c.JSON(200, tweet)
}

func GetTweet(c *gin.Context) {
	id := c.Param("id")

	var tweet model.Tweet
	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// Validate the tweet ID
	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.JSON(400, gin.H{"error": "Invalid tweet ID"})
		return
	}

	// Fetch the tweet with its likes count
	result := db.Preload("Likes").First(&tweet, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Tweet not found"})
		} else {
			c.JSON(500, gin.H{"error": "Internal server error"})
		}
		return
	}

	// Prepare the response
	c.JSON(200, gin.H{
		"tweet": tweet,
	})
}

func GetTweets(c *gin.Context) {
	var tweets []model.Tweet
	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	result := db.Find(&tweets)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, tweets)
}

func GetTweetsByUserID(c *gin.Context) {
	userID := c.Param("id")

	var tweets []model.Tweet
	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("user_id = ?", userID).Find(&tweets)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, tweets)
}
