package controller


import (
	"github.com/gin-gonic/gin"
	"github.com/ashsajal1/go-tweet/model"
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
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Fetch the tweet
    result := db.First(&tweet, id)
    if result.Error != nil {
        c.JSON(404, gin.H{"error": "Tweet not found"})
        return
    }

    // Get the likes count for the tweet
    var likeCount int64
    db.Model(&model.Like{}).Where("tweet_id = ?", tweet.ID).Count(&likeCount)

    // Prepare the response
    c.JSON(200, gin.H{
        "tweet":      tweet,
        "likes_count": likeCount,
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
