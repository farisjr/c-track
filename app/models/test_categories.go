package models

import "gorm.io/gorm"

type TestCategories struct {
	gorm.Model
	ID                   int    `gorm:"primaryKey" json:"id" form:"id"`
	Test_categories_Name string `json:"test_categories_name" form:"test_categories_name"`
	//Tests                []Tests `gorm:"foreignKey:TestCategoriesID"`
}
