package database

import (
	"app/config"
	"app/models"
)

func CreateChecker(checker models.Checker) (interface{}, error) {
	if err := config.DB.Save(&checker).Error; err != nil {
		return nil, err
	}
	return checker, nil
}

func GetCheckers() (interface{}, error) {
	var checkers []models.Checker
	if err := config.DB.Find(&checkers).Error; err != nil {
		return nil, err
	}
	return checkers, nil
}

func GetCheckerById(id int) (models.Checker, error) {
	var checker models.Checker
	if err := config.DB.Find(&checker, "employee_id=?", id).Error; err != nil {
		return checker, err
	}
	return checker, nil
}
