package models

import (
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Name                   string `json:"name"`
	MedicalFacilityName    string `json:"medical_facility_name"`
	MedicalFacilityAddress string `json:"medical_facility_address"`
	User                   User   `gorm:"foreignKey:ID" json:"user_id"`
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
