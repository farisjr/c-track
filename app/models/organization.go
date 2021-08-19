package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Organization_ID      int    `gorm:"primerykey;AUTO_INCREMENT" json:"id"`
	Organization_Name    string `gorm:"type:varchar(100);not null" json:"name"`
	Organization_Address string `gorm:"type:varchar(100);not null" json:"address"`
	Organization_City    string `gorm:"type:varchar(16);not null" json:"city"`
	Organization_Prov    string `gorm:"type:varchar(16);not null" json:"prov"`
}
