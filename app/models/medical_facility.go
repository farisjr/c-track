package models

import "gorm.io/gorm"

type Medical_Facility struct {
	gorm.Model
	Faskes_ID      string `gorm:"primerykey;type:varchar(30);not null" json:"faskesid"`
	Faskes_Name    string `gorm:"type:varchar(100);not null" json:"faskesname"`
	Faskes_Address string `gorm:"type:varchar(100);not null" json:"faskesadrress"`
	Faskes_City    string `gorm:"type:varchar(100);not null" json:"faskescity"`
	Faskes_Prov    string `gorm:"type:varchar(100);not null" json:"province"`
}
