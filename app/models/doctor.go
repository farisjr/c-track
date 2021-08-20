package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	PracticeLicenseID      int    `gorm:"primaryKey" json:"practice_license_id" form:"practice_license_id"`
	Name                   string `gorm:"type:varchar(45);not null" json:"name"`
	MedicalFacilityName    string `gorm:"type:varchar(45);not null" json:"medical_facility_name"`
	MedicalFacilityAddress string `gorm:"type:varchar(45);not null" json:"medical_facility_address"`
	// UserID                 string `gorm:"type:char(30);not null" json:"user_id"`
	// User                   User   `gorm:"foreignKey:ID"`
}
