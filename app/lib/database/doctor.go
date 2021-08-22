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
	var doctor []models.Doctor
	if err := config.DB.Find(&doctor).Error; err != nil {
		return nil, err
	}
	return doctor, nil
}

func GetDoctorById(id int) (models.Doctor, error) {
	var doctor models.Doctor
	if err := config.DB.Find(&doctor, "id=?", id).Error; err != nil {
		return doctor, err
	}
	return doctor, nil
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
	var doctor models.Doctor
	config.DB.Find(&doctor, "id=?", id)
	return doctor
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
