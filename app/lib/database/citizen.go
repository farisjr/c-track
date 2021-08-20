package database

import (
	"app/config"
	"app/models"
)

func CreatePatient(patient models.Patient) (interface{}, error) {
	if err := config.DB.Save(&patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func GetPatientByNIK(NIK string) (models.Patient, error) {
	var patient models.Patient
	var empty models.Patient

	if err := config.DB.Where("Patient_ID = ?", NIK).First(&patient).Error; err != nil {
		return empty, err
	}
	if err := config.DB.Find(&patient, NIK).Error; err != nil {
		return empty, err
	}
	return patient, nil
}

func UpdatePatient(Patient_ID int, patient interface{}) (interface{}, error) {
	if err := config.DB.Find(&patient, "Patient_ID=?", Patient_ID).Save(&patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func DeleteCitizen(Patient_ID int) (interface{}, error) {
	var patient models.Patient
	if err := config.DB.Find(&patient, "Patient_ID=?", Patient_ID).Delete(&patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}
