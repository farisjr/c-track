package models

import "gorm.io/gorm"

type Test_Category struct {
	gorm.Model
	Category_ID   string `gorm:"type:char(16);not null" json:"categoryid"`
	Category_Name string `gorm:"type:varchar(100);not null" json:"categoryname"`
}
