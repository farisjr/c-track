package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Practice_License_ID      int    `gorm:"primerykey;unique;not null" json:"practice_license_id"`
	Name                     string `gorm:"type:varchar(45);not null" json:"name"`
	Medical_Facility_Name    string `gorm:"type:varchar(45);not null" json:"medical_facility_name"`
	Medical_Facility_Address string `gorm:"type:varchar(45);not null" json:"medical_facility_address"`
	User_ID                  string `gorm:"type:char(30);not null" json:"user_id"`
	User                     User   `gorm:"foreignkey:ID"`
}
