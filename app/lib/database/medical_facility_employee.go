package database

import (
	"app/config"
	"app/models"
)

func CreateMedicalFacilityEmployee(medical_facility_employee models.Medical_Facility_Employee) (interface{}, error) {
	if err := config.DB.Save(&medical_facility_employee).Error; err != nil {
		return nil, err
	}
	return medical_facility_employee, nil
}
