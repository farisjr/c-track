package models

import "gorm.io/gorm"

type TestCategories struct {
	gorm.Model
	TestCategoriesName string `json:"test_categories_name" form:"test_categories_name"`
}
