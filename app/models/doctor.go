package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	PracticeLicenseID      int    `gorm:"primaryKey;unique;not null" json:"practice_license_id" form:"practice_license_id"`
	UserID                 int    `gorm:"primarykey;unique;not null" json:"user_id"`
	Name                   string `gorm:"type:varchar(45);not null" json:"name"`
	MedicalFacilityName    string `gorm:"type:varchar(45);not null" json:"medical_facility_name"`
	MedicalFacilityAddress string `gorm:"type:varchar(45);not null" json:"medical_facility_address"`
	User                   User   `gorm:"foreignKey:UserID"`
}
