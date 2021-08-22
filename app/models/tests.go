package models

import (
	"gorm.io/gorm"
)

type Tests struct {
	gorm.Model
	Result string `json:"result" form:"result"`

	// 1 to many with test categories
	TestCategories TestCategories `gorm:"foreignKey:ID;" json:"category_id"`
	// 1 to many with patient
	Patient Patient `gorm:"foreignKey:ID" json:"patient_id"`
}
