package database

import (
	"app/config"
	"app/models"
)

func CreateTestCategories(testCategories models.TestCategories) (models.TestCategories, error) {
	if err := config.DB.Save(&testCategories).Error; err != nil {
		return testCategories, err
	}
	return testCategories, nil
}

func GetTestCategories() (interface{}, error) {
	var testCategories []models.TestCategories
	if err := config.DB.Find(&testCategories).Error; err != nil {
		return nil, err
	}
	return testCategories, nil
}

func GetTestCategory(id int) (models.TestCategories, error) {
	var testCategories models.TestCategories

	if err := config.DB.Find(&testCategories, "id=?", id).Error; err != nil {
		return testCategories, err
	}
	return testCategories, nil
}

func DeleteTestCategory(id int) (interface{}, error) {
	var testCategories []models.TestCategories
	if err := config.DB.Find(&testCategories, "id=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&testCategories, "id=?", id).Error; err != nil {
		return nil, err
	}
	return testCategories, nil
}

func UpdateTestCategory(testCategories models.TestCategories) (models.TestCategories, error) {
	if err := config.DB.Save(&testCategories).Error; err != nil {
		return testCategories, err
	}

	return testCategories, nil
}
