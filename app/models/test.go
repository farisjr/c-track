package models

import (
	"time"

	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Test_ID       int    `gorm:"primerykey;not null" json:"id"`
	NIK           string `gorm:"type:char(16);not null" json:"nik"`
	TestDate      time.Time
	CategoryID    string `gorm:"type:char(16);not null" json:"categoryid"`
	StrID         string `gorm:"type:char(16);unique;not null" json:"strid"`
	Result_Date   time.Time
	Result_Status bool
	Note          string        `gorm:"type:varchar(255);not null" json:"note"`
	Category      Test_Category `gorm:"foreignkey:ID"`
	Str           Doctor        `gorm:"foreignkey:ID"`
}
