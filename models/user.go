package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       		string `json:"name"`
	Email      		string `json:"email" gorm:"unique"`
	Password   		string `json:"-"`
	Confirmed  		bool   `json:"confirmed" gorm:"default:false"`
	ConfirmToken 	string `json:"-"`
	Blogs    []Blog `json:"Blogs" gorm:"foreignKey:UserID"`
}