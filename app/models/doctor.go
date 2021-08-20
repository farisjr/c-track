package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Practice_License_ID      int    `gorm:"primerykey;unique;not null" json:"strid"`
	Name                     string `gorm:"type:varchar(45);not null" json:"nik"`
	Medical_Facility_Name    string `gorm:"type:varchar(45);not null" json:"faskesid"`
	Medical_Facility_Address string `gorm:"type:varchar(45);not null" json:"email"`
	User_ID                  string `gorm:"type:char(30);not null" json:"password"`
	User                     User   `gorm:"foreignkey:ID"`
}
