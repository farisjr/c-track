package database

import (
	"app/config"
	"app/models"
)

func CreateDoctors(addDoctor models.Doctor) (interface{}, error) {
	if err := config.DB.Save(&addDoctor).Error; err != nil {
		return nil, err
	}
	return addDoctor, nil
}

func GetDoctors() (interface{}, error) {
	var doctors []models.Doctor
	if err := config.DB.Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

func GetDoctorsById(id int) (models.Doctor, error) {
	var doctors models.Doctor
	if err := config.DB.Find(&doctors, "id=?", id).Error; err != nil {
		return doctors, err
	}
	return doctors, nil
}

func DeleteDoctorsById(deleteDoctor models.Doctor) (interface{}, error) {

	if err := config.DB.Delete(&deleteDoctor).Error; err != nil {
		return nil, err
	}
	return deleteDoctor, nil
}

//update test info from database
func UpdateDoctors(updateDoctors models.Doctor) (interface{}, error) {
	if tx := config.DB.Save(&updateDoctors).Error; tx != nil {
		return updateDoctors, tx
	}
	return updateDoctors, nil
}

//get 1 specified test with test struct output
func GetUpdateDoctors(id int) models.Doctor {
	var doctors models.Doctor
	config.DB.Find(&doctors, "id=?", id)
	return doctors
}
