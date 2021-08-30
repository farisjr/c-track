package models

import (
	"database/sql"
	"time"
)

type Doctor struct {
	DoctorID               int    `gorm:"primaryKey; unique; not null" json:"doctor_id" form:"doctor_id"`
	UserID                 int    `json:"user_id" form:"user_id"`
	Name                   string `json:"name"`
	MedicalFacilityName    string `json:"medical_facility_name"`
	MedicalFacilityAddress string `json:"medical_facility_address"`
	User                   User   `gorm:"foreignKey:UserID"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              sql.NullTime `gorm:"index"`
}

type DoctorResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    []Doctor `json:"data"`
}

type DoctorResponseSingle struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Doctor `json:"data"`
}
