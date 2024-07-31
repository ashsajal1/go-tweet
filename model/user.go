package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Tweets   []Tweet `gorm:"foreignkey:UserID"`
	Likes    []Like  `gorm:"foreignkey:UserID"`
}
