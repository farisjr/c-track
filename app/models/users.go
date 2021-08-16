package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID       int    `gorm:"primaryKey" json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
	Gender   string `json:"gender" form:"gender"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}
