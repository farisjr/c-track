package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	PatientID      int            `gorm:primaryKey"; not null" json:"patient_id"`
	Dob            time.Time      `gorm:"; not null" json:"borndate"`
	Pob            string         `gorm:"type:varchar(30); not null" json:"bornplace"`
	Address        string         `gorm:"type:varchar(100); not null" json:"address"`
	City           string         `gorm:"type:varchar(100); not null" json:"city"`
	Province       string         `gorm:"type:varchar(100); not null" json:"province"`
	Gender         Gender         `gorm:"not null" json:"gender"`
	Blood_type     Blood_type     `gorm:"not null" json:"blood type"`
	Religion       Religion       `gorm:"not null" json:"religion"`
	Marital_Status Marital_Status `gorm:"not null" json:"marital status"`
	User           User           `gorm:"foreignKey:ID"`
}

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
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

type Marital_Status string

const (
	Mariage Marital_Status = "Mariage"
	Single  Marital_Status = "Single"
)
