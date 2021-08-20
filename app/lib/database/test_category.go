package database

import (
	"app/config"
	"app/models"
)

func CreateTestCategory(test_category models.Test_Category) (interface{}, error) {
	if err := config.DB.Save(&test_category).Error; err != nil {
		return nil, err
	}
	return test_category, nil
}
