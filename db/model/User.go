package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string
}
