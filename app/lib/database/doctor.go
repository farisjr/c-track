package database

import (
	"app/config"
	"app/models"
)

func CreateDoctor(addDoctor models.Doctor) (interface{}, error) {
	if err := config.DB.Save(&addDoctor).Error; err != nil {
		return nil, err
	}
	return addDoctor, nil
}

func GetDoctor() (interface{}, error) {
	var doctors []models.Doctor
	if err := config.DB.Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

func GetDoctorById(id int) (models.Doctor, error) {
	var doctors models.Doctor
	if err := config.DB.Find(&doctors, "id=?", id).Error; err != nil {
		return doctors, err
	}
	return doctors, nil
}

func DeleteDoctorById(deleteDoctor models.Doctor) (interface{}, error) {

	if err := config.DB.Delete(&deleteDoctor).Error; err != nil {
		return nil, err
	}
	return deleteDoctor, nil
}

//update test info from database
func UpdateDoctor(updateDoctors models.Doctor) (interface{}, error) {
	if tx := config.DB.Save(&updateDoctors).Error; tx != nil {
		return updateDoctors, tx
	}
	return updateDoctors, nil
}

//get 1 specified test with test struct output
func GetUpdateDoctor(id int) models.Doctor {
	var doctors models.Doctor
	config.DB.Find(&doctors, "id=?", id)
	return doctors
}

//Login for doctor with matching username and password
func LoginDoctorDB(username, password string) (models.Doctor, error) {
	var doctor models.Doctor
	var err error
	if err = config.DB.Where("username=? AND password=?", username, password).First(&doctor).Error; err != nil {
		return doctor, err
	}
	if err := config.DB.Save(doctor).Error; err != nil {
		return doctor, err
	}
	return doctor, err
}
