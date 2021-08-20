package models

import "gorm.io/gorm"

type Checker struct {
	gorm.Model
	Employee_ID    int    `gorm:"primerykey;not null" json:"employeeid"`
	Name           string `gorm:"type:varchar(45);not null" json:"organizationid"`
	Office_Name    string `gorm:"type:varchar(45);;not null" json:"employeeid"`
	Office_Address string `gorm:"type:varchar(30);not null" json:"email"`
	User_ID        string `gorm:"type:varchar(30);not null" json:"password"`
	User           User   `gorm:"foreignkey:User_ID"`
}
