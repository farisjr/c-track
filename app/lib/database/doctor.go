package database

import (
	"app/config"
	"app/models"
)

func CreateDoctor(doctor models.Doctor) (interface{}, error) {
	if err := config.DB.Save(&doctor).Error; err != nil {
		return nil, err
	}
	return doctor, nil
}
