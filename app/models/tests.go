package models

import (
	"time"

	"gorm.io/gorm"
)

type Tests struct {
	gorm.Model
	TestID           int `gorm:"primaryKey;unique;not null" json:"test_id" form:"test_id"`
	TestDate         time.Time
	ResultDate       time.Time
	Result           bool
	PatientID        int `gorm:"primarykey;unique;not null" json:"patient_id"`
	TestCategoriesID int `gorm:"primarykey;unique;not null" json:"test_categories_id"`

	// 1 to many with test categories
	TestCategories TestCategories `gorm:"foreignkey:TestCategoriesID"`
	// 1 to many with patient
	Patient Patient `gorm:"foreignkey:PatientID"`
}
