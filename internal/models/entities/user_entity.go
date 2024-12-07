package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required" gorm:"unique;size:100"`
	UserName string `json:"user_name" binding:"required" gorm:"unique;size:100"`
	Password string `json:"password" binding:"required" gorm:"not null"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}
