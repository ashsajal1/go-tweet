package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ashsajal1/go-tweet/model"
)

func CreateLike(c *gin.Context) {
	var like model.Like
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&like)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, like)
}

func GetLike(c *gin.Context) {
	id := c.Param("id")

	var like model.Like
	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	result := db.First(&like, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Like not found"})
		return
	}

	c.JSON(200, like)
}

func DeleteLike(c *gin.Context) {
	id := c.Param("id")

	var like model.Like
	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	result := db.First(&like, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Like not found"})
		return
	}

	result = db.Delete(&like)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "successfully deleted"})
}
