package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NIK          string `gorm:"primerykey;type:char(16);unique;not null" json:"nik"`
	Email        string `gorm:"type:char(16);not null" json:"email"`
	Password     string `gorm:"type:varchar(100);not null" json:"password"`
	Phone_number string `gorm:"type:char(16);not null" json:"phone number"`
}
