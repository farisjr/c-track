package models

import "gorm.io/gorm"

type Checker struct {
	gorm.Model
	Employee_ID    int    `gorm:"primarykey;not null" json:"employee_id"`
	Name           string `gorm:"type:varchar(45);not null" json:"organization_id"`
	Office_Name    string `gorm:"type:varchar(45);;not null" json:"office_name"`
	Office_Address string `gorm:"type:varchar(30);not null" json:"office_address"`
	User           User   `gorm:"foreignkey:User_ID"`
}
