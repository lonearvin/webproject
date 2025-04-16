package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" binding:"required" gorm:"unique"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `gorm:"unique"`
}
