package database

import (
	"app/config"
	"app/models"
)

func CreateTestCategories(testCategories models.TestCategories) (interface{}, error) {
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

func GetTestCategoriesId(id int) (models.TestCategories, error) {
	var testCategories models.TestCategories
	if err := config.DB.Find(&testCategories, "id=?", id).Error; err != nil {
		return testCategories, err
	}
	return testCategories, nil
}

func DeleteTestCategoriesById(testCategories models.TestCategories) (interface{}, error) {
	if err := config.DB.Delete(&testCategories).Error; err != nil {
		return nil, err
	}
	return testCategories, nil
}

//update test categories info from database
func UpdateTestCategories(testCategories models.TestCategories) (interface{}, error) {
	if tx := config.DB.Save(&testCategories).Error; tx != nil {
		return testCategories, tx
	}
	return testCategories, nil
}

//get 1 specified test categories with testCategories struct output
func GetUpdateTestCategories(id int) models.TestCategories {
	var testCategories models.TestCategories
	config.DB.Find(&testCategories, "id=?", id)
	return testCategories
}
