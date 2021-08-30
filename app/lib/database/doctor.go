package database

import (
	"app/config"
	"app/middlewares"
	"app/models"
)

func CreateDoctor(doctor models.User) (models.User, error) {
	if err := config.DB.Save(&doctor).Error; err != nil {
		return doctor, err
	}
	return doctor, nil
}

func GetDoctor() (models.Doctor, error) {
	var doctor models.Doctor
	if err := config.DB.Find(&doctor).Error; err != nil {
		return doctor, err
	}
	return doctor, nil
}

func GetOneDoctor(id int) (models.User, error) {
	var doctor models.User
	// var count int64
	// if err1 := config.DB.Model(&doctor).Where("user_id=?", id).Count(&count).Error; count == 0 {
	// 	return doctor, err1
	// }
	if err := config.DB.Find(&doctor, "user_id=?", id).Error; err != nil {
		return doctor, err
	}
	return doctor, nil
}

//update doctor from database
func UpdateDoctor(doctor models.User) (models.User, error) {
	if tx := config.DB.Save(&doctor).Error; tx != nil {
		return doctor, tx
	}
	return doctor, nil
}

//Login for doctor with matching user id and password
func DoctorLoginDB(userId int, password string) (models.User, error) {
	var doctor models.User
	var err error
	if err = config.DB.Where("user_id=? AND password=?", userId, password).First(&doctor).Error; err != nil {
		return doctor, err
	}
	doctor.Token, err = middlewares.CreateDoctorToken(int(doctor.UserID))
	if err != nil {
		return doctor, err
	}
	if err := config.DB.Save(doctor).Error; err != nil {
		return doctor, err
	}
	return doctor, err
}
