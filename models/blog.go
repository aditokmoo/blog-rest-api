package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"User" gorm:"foreignKey:UserID"`
}