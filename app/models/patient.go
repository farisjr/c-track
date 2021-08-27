package models

import (
	"time"
)

type Patient struct {
	PatientID      int            `gorm:"primaryKey; unique; not null" json:"patient_id"`
	UserID         int            `json:"user_id" form:"user_id"`
	Fullname       string         `gorm:"type:varchar(30); not null" json:"fullname"`
	Dob            time.Time      `gorm:"not null" json:"borndate"`
	Pob            string         `gorm:"type:varchar(30); not null" json:"bornplace"`
	Address        string         `gorm:"type:varchar(100); not null" json:"address"`
	City           string         `gorm:"type:varchar(100); not null" json:"city"`
	Province       string         `gorm:"type:varchar(100); not null" json:"province"`
	Gender         Gender         `gorm:"type:varchar(10);not null" json:"gender"`
	Blood_type     Blood_type     `gorm:"type:varchar(2);not null" json:"blood_type"`
	Religion       Religion       `gorm:"type:varchar(10);not null" json:"religion"`
	Marital_Status Marital_Status `gorm:"type:varchar(15);not null" json:"marital_status"`
	User           User           `gorm:"foreignKey:UserID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time `gorm:"index"`
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
