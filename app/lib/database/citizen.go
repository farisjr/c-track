package database

import (
	"app/config"
	"app/models"
)

func CreateCitizen(citizen models.Citizen) (interface{}, error) {
	if err := config.DB.Save(&citizen).Error; err != nil {
		return nil, err
	}
	return citizen, nil
}

func GetCitizenByNIK(NIK string) (models.Citizen, error) {
	var citizen models.Citizen
	var empty models.Citizen

	if err := config.DB.Where("NIK = ?", NIK).First(&citizen).Error; err != nil {
		return empty, err
	}
	if err := config.DB.Find(&citizen, NIK).Error; err != nil {
		return empty, err
	}
	return citizen, nil
}

func GetCitizenByName(name string) (models.Citizen, error) {
	var citizen models.Citizen
	if e := config.DB.Where("Name = ?", name).Find(&citizen).Error; e != nil {
		return citizen, e
	}

	return citizen, nil
}
