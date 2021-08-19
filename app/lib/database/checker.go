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
