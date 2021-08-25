package models

import (
	"database/sql"
	"time"
)

type Checker struct {
	EmployeeID    int    `gorm:"primarykey;unique;not null" json:"employee_id"`
	UserID        int    `json:"user_id" form:"user_id"`
	Name          string `gorm:"type:varchar(45);not null" json:"name"`
	OfficeName    string `gorm:"type:varchar(45);not null" json:"office_name"`
	OfficeAddress string `gorm:"type:varchar(30);not null" json:"office_address"`
	User          User   `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
}
