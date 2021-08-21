package models

import "gorm.io/gorm"

type Checker struct {
	gorm.Model
	UserID        int    `gorm:"primarykey;unique;not null" json:"user_id"`
	Name          string `gorm:"type:varchar(45);not null" json:"name"`
	OfficeName    string `gorm:"type:varchar(45);not null" json:"office_name"`
	OfficeAddress string `gorm:"type:varchar(30);not null" json:"office_address"`
	User          User   `gorm:"foreignKey:ID"`
}
