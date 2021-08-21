package database

import (
	"app/config"
	"app/models"
	"errors"
)

func CreatePatient(patient models.Patient) (interface{}, error) {
	if err := config.DB.Save(&patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func GetPatient() (interface{}, error) {
	var patients []models.Patient
	if err := config.DB.Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

func GetPatientById(id int) (models.Patient, error) {
	var patient models.Patient
	var empty models.Patient

	if err := config.DB.Find(&patient, "id=?", id).Error; err != nil {
		return empty, err
	}
	return patient, nil
}

func UpdatePatient(patient models.Patient) (interface{}, error) {
	if err := config.DB.Save(&patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func DeletePatient(id int) (interface{}, error) {
	var patient models.Patient
	if err := config.DB.Find(&patient, "id=?", id).Delete(&patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func CheckDuplicatePatient(patient models.Patient) error {
	var patients models.Patient
	err := config.DB.Where("id = ?", patients.ID).First(&patient).Error
	if err != nil {
		return nil
	}
	err = errors.New("id Already Exist")
	return err
}

//Login for patient with matching username and password
func PatientLoginDB(username, password string) (models.Patient, error) {
	var patient models.Patient
	var err error
	if err = config.DB.Where("username=? AND password=?", username, password).First(&patient).Error; err != nil {
		return patient, err
	}
	if err := config.DB.Save(patient).Error; err != nil {
		return patient, err
	}
	return patient, err
}
