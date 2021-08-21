package models

import "gorm.io/gorm"

type Checker struct {
	gorm.Model
	EmployeeID    int    `gorm:"primarykey;not null" json:"employee_id"`
	Name           string `gorm:"type:varchar(45);not null" json:"organization_id"`
	OfficeName    string `gorm:"type:varchar(45);;not null" json:"office_name"`
	OfficeAddress string `gorm:"type:varchar(30);not null" json:"office_address"`
	User           User   `gorm:"foreignkey:UserID"`
}
