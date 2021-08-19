package models

import "gorm.io/gorm"

type Citizen struct {
	gorm.Model
	NIK            int            `gorm:"primerykey;not null" json:"nik"`
	Name           string         `gorm:"type:varchar(100);not null" json:"email"`
	Dob_date       string         `gorm:"type:varchar(30);not null" json:"borndate"`
	Dob_place      string         `gorm:"type:varchar(30);not null" json:"bornplace"`
	Address        string         `gorm:"type:varchar(100);not null" json:"address"`
	City           string         `gorm:"type:varchar(100);not null" json:"city"`
	Province       string         `gorm:"type:varchar(100);not null" json:"province"`
	Gender         Gender         `gorm:"not null" json:"gender"`
	Blood_type     Blood_type     `gorm:"not null" json:"blood type"`
	Religion       Religion       `gorm:"not null" json:"religion"`
	Mariage_status Mariage_Status `gorm:"not null" json:"mariage status"`
}

type Gender string

const (
	Pria   Gender = "Pria"
	Wanita Gender = "Wanita"
)

type Blood_type string

const (
	A  Blood_type = "A"
	B  Blood_type = "B"
	AB Blood_type = "AB"
	O  Blood_type = "O"
)

type Religion string

const (
	Islam    Religion = "Islam"
	Kristen  Religion = "Kristen"
	Hindu    Religion = "Hindu"
	Budha    Religion = "Budha"
	Katholik Religion = "katholik"
)

type Mariage_Status string

const (
	Kawin      Mariage_Status = "Kawin"
	TidakKawin Mariage_Status = "Tidak Kawin"
)
