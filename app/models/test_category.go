package models

import "gorm.io/gorm"

type Test_Category struct {
	gorm.Model
	Test_Category_ID int    `gorm:"primerykey;not null" json:"test_categoryid"`
	Category_Name    string `gorm:"type:varchar(45);not null" json:"categoryname"`
}
