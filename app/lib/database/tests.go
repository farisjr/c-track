package database

import (
	"app/config"
	"app/models"
)

func CreateTests(tests models.Tests) (models.Tests, error) {
	if err := config.DB.Save(&tests).Error; err != nil {
		return tests, err
	}
	return tests, nil
}

func GetAllTests() (interface{}, error) {
	var tests []models.Tests
	if err := config.DB.Find(&tests).Error; err != nil {
		return nil, err
	}
	return tests, nil
}

func GetTestsId(id int) (interface{}, error) {
	var tests models.Tests
	var count int64
	if err1 := config.DB.Model(&tests).Where("id=?", id).Count(&count).Error; count == 0 {
		return tests, err1
	}
	if err := config.DB.Find(&tests, "id=?", id).Error; err != nil {
		return tests, err
	}
	return tests, nil
}

func DeleteTestsById(id int) (interface{}, error) {
	var tests models.Tests
	if err := config.DB.Where("id=?", id).Delete(&tests).Error; err != nil {
		return nil, err
	}
	return tests, nil
}

//update test info from database
func UpdateTests(tests models.Tests) (models.Tests, error) {
	if tx := config.DB.Save(&tests).Error; tx != nil {
		return tests, tx
	}
	return tests, nil
}

//get 1 specified test with test struct output
func GetUpdateTests(id int) models.Tests {
	var tests models.Tests
	config.DB.Find(&tests, "id=?", id)
	return tests
}
