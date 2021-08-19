package database

import (
	"app/config"
	"app/models"
)

func CreateMedicalFacility(medical_facility models.Medical_Facility) (interface{}, error) {
	if err := config.DB.Save(&medical_facility).Error; err != nil {
		return nil, err
	}
	return medical_facility, nil
}
