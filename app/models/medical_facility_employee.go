package models

import "gorm.io/gorm"

type Medical_Facility_Employee struct {
	gorm.Model
	Str_ID    string           `gorm:"type:char(16);unique;not null" json:"strid"`
	Faskes_ID string           `gorm:"type:char(16);not null" json:"faskesid"`
	Nakes_ID  string           `gorm:"type:char(16);not null" json:"nakesid"`
	Str       Doctor           `gorm:"foreignkey:ID"`
	Faskes    Medical_Facility `gorm:"foreignkey:ID"`
	Nakes     Doctor           `gorm:"foreignkey:ID"`
}
