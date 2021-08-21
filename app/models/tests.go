package models

import (
	"time"

	"gorm.io/gorm"
)

type Tests struct {
	gorm.Model
	TestDate         time.Time
	ResultDate       time.Time
	Result           bool
	PatientID        int `gorm:"primaryKey; unique; not null" json:"patient_id"`
	TestCategoriesID int `gorm:"primaryKey; unique; not null" json:"test_categories_id"`

	// 1 to many with test categories
	TestCategories TestCategories `gorm:"foreignKey:ID"`
	// 1 to many with patient
	Patient Patient `gorm:"foreignKey:ID"`
}
