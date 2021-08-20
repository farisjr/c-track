package models

import (
	"gorm.io/gorm"
)

type Tests struct {
	gorm.Model
	//ID int `gorm:"primaryKey" json:"id" form:"id"`
	//Test_Date     time.Time
	//Result_Date   time.Time
	Result bool

	// 1 to many with test categories
	//TestCategoriesID int `json:"test_categories_id" form:"test_categories_id"`

	//Patient_ID         int            `gorm:"not null" json:"patient_id"`

}
