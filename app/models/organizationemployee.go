package models

import "gorm.io/gorm"

type Organization_Employee struct {
	gorm.Model
	Employee_ID     int          `gorm:"primerykey;not null" json:"employeeid"`
	Organization_ID string       `gorm:"type:char(16);not null" json:"organizationid"`
	NIK             string       `gorm:"type:char(16);not null" json:"nik"`
	Organization    Organization `gorm:"foreignkey:ID" json:"-"`
}
