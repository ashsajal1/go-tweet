package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID   uint
	TweetID  uint
	Disliked bool
}