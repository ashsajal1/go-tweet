package controller

import (
	"log"

	"github.com/ashsajal1/go-tweet/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, user)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user model.User
	db, err := model.GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	result := db.First(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func GetUsers(c *gin.Context) {
	var users []model.User
	db, err := model.GetDB()
	if err != nil {
		log.Println("Error getting database connection:", err)
		c.JSON(500, gin.H{"error": "Failed to get database connection"})
		return
	}

	result := db.Find(&users)
	if result.Error != nil {
		log.Println("Error querying users:", result.Error)
		c.JSON(500, gin.H{"error": "Failed to query users"})
		return
	}

	log.Println("Retrieved users:", users)
	c.JSON(200, users)
}
