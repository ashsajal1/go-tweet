package model

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	UserID  uint
	Content string
	Likes   []Like `gorm:"foreignkey:TweetID"`
}
