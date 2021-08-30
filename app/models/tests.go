package models

import (
	"database/sql"
	"time"
)

type Tests struct {
	TestID           int    `gorm:"primaryKey; unique; not null" json:"test_id" form:"test_id"`
	PatientID        int    `json:"patient_id" form:"patient_id"`
	TestCategoriesID int    `json:"testcategories_id" form:"testcategories_id"`
	DoctorID         int    `json:"doctor_id" form:"doctor_id"`
	Result           string `gorm:"type:varchar(100);" json:"result" form:"result"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime `gorm:"index"`
}
