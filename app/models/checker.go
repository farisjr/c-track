package models

import "gorm.io/gorm"

type Checker struct {
	gorm.Model
	EmployeeID    int    `gorm:"primarykey;unique;not null" json:"employeeid"`
	Name          string `gorm:"type:varchar(45);not null" json:"name"`
	OfficeName    string `gorm:"type:varchar(45);not null" json:"office_name"`
	OfficeAddress string `gorm:"type:varchar(30);not null" json:"office_address"`
	User          User   `gorm:"foreignKey:ID"`
}
