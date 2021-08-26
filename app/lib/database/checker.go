package database

import (
	"app/config"
	"app/models"
)

func CreateChecker(checker models.Checker) (models.Checker, error) {
	if err := config.DB.Save(&checker).Error; err != nil {
		return checker, err
	}
	return checker, nil
}

func GetCheckers() (models.Checker, error) {
	var checkers models.Checker
	if err := config.DB.Find(&checkers).Error; err != nil {
		return checkers, err
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

func UpdateChecker(checker models.Checker) (models.Checker, error) {
	if err := config.DB.Save(&checker).Error; err != nil {
		return checker, err
	}
	return checker, nil
}

/*
//Login for checker with matching email and password
func LoginCheckerDB(username, password string) (models.Checker, error) {
	var checker models.Checker
	var err error
	if err = config.DB.Where("username=? AND password=?", username, password).First(&checker).Error; err != nil {
		return checker, err
	}
	if err := config.DB.Save(checker).Error; err != nil {
		return checker, err
	}
	return checker, err
}
*/
