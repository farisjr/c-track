package models

import (
	"database/sql"
	"time"
)

type Tests struct {
<<<<<<< HEAD
	TestsID          int `gorm:"primaryKey; unique; not null" json:"test_id"`
=======
	TestsID          int `gorm:"primaryKey; not null" json:"test_id" form:"test_id"`
>>>>>>> development
	PatientID        int `json:"patient_id" form:"patient_id"`
	TestCategoriesID int `json:"testcategories_id" form:"testcategories_id"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime `gorm:"index"`
	// 1 to many with test categories
	TestCategories TestCategories `gorm:"foreignKey:TestCategoriesID"`
	// 1 to many with patient
	Patient Patient `gorm:"foreignKey:PatientID"`
	Result  string  `gorm:"type:varchar(100);" json:"result" form:"result"`
}
