package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Str_ID       string           `gorm:"primerykey;type:char(16);unique;not null" json:"strid"`
	NIK          string           `gorm:"type:char(16);not null" json:"nik"`
	Faskes_ID    string           `gorm:"type:char(16);not null" json:"faskesid"`
	Email        string           `gorm:"type:char(16);not null" json:"email"`
	Password     string           `gorm:"type:char(30);not null" json:"password"`
	Phone_Number string           `gorm:"type:char(16);not null" json:"phonenumber"`
	Faskes       Medical_Facility `gorm:"foreignkey:ID"`
}
