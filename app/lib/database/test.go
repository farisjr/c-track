package database

import (
	"app/config"
	"app/models"
)

func CreateTest(test models.Test) (interface{}, error) {
	if err := config.DB.Save(&test).Error; err != nil {
		return nil, err
	}
	return test, nil
}
