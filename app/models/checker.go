package models

import "gorm.io/gorm"

type Checker struct {
	gorm.Model
	NIK             int                   `gorm:"primerykey;not null" json:"NIK"`
	Organization_ID string                `gorm:"type:char(16);not null" json:"organizationid"`
	Employee_ID     string                `gorm:"type:char(16);unique;not null" json:"employeeid"`
	Email           string                `gorm:"type:varchar(30);not null" json:"email"`
	Password        string                `gorm:"type:varchar(30);not null" json:"password"`
	Phone_Number    string                `gorm:"type:char(16);not null" json:"phone number"`
	Organization    Organization          `gorm:"foreignkey:ID"`
	Employee        Organization_Employee `gorm:"foreignkey:ID"`
}
