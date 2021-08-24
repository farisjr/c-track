package models

import (
	"gorm.io/gorm"
)

type TestCategories struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}

type TestCategoriesResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    []TestCategories `json:"data"`
}

type TestCategoriesResponseSingle struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    TestCategories `json:"data"`
}
