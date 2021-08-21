package models

import "gorm.io/gorm"

type TestCategories struct {
	gorm.Model
	TestCategoriesID   int    `gorm:"primaryKey;unique;not null" json:"test_categories_id" form:"test_categories_id"`
	TestCategoriesName string `json:"test_categories_name" form:"test_categories_name"`
	//Many to one with test
	Tests []Tests `gorm:"foreignKey:TestCategoriesID"`
}
