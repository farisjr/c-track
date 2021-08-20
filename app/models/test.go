package models

import (
	"time"

	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Test_ID            int `gorm:"primerykey;not null" json:"test_id"`
	Test_Date          time.Time
	Result_Date        time.Time
	Result_Status      bool
	Test_Categories_ID int            `gorm:"type:char(16);not null" json:"test_categories_id"`
	Patient_ID         int            `gorm:"not null" json:"patient_id"`
	TestCategories     TestCategories `gorm:"foreignkey:ID"`
	Patient            Patient        `gorm:"foreignkey:ID"`
}
