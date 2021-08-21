package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	UserID                 int    `gorm:"primaryKey; unique; not null" json:"user_id"`
	Name                   string `gorm:"type:varchar(45); not null" json:"name"`
	MedicalFacilityName    string `gorm:"type:varchar(45); not null" json:"medical_facility_name"`
	MedicalFacilityAddress string `gorm:"type:varchar(45); not null" json:"medical_facility_address"`
	User                   User   `gorm:"foreignKey:ID"`
}
