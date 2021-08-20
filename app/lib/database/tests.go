package database

import (
	"app/config"
	"app/models"
)

func CreateTests(createTests models.Tests) (interface{}, error) {
	if err := config.DB.Save(&createTests).Error; err != nil {
		return nil, err
	}
	return createTests, nil
}

func GetTests() (interface{}, error) {
	var tests []models.Tests
	if err := config.DB.Find(&tests).Error; err != nil {
		return nil, err
	}
	return tests, nil
}

func GetTestsById(id int) (models.Tests, error) {
	var tests models.Tests
	if err := config.DB.Find(&tests, "id=?", id).Error; err != nil {
		return tests, err
	}
	return tests, nil
}

func DeleteTestsById(deleteTest models.Tests) (interface{}, error) {

	if err := config.DB.Delete(&deleteTest).Error; err != nil {
		return nil, err
	}
	return deleteTest, nil
}

//update test info from database
func UpdateTests(updateTests models.Tests) (interface{}, error) {
	if tx := config.DB.Save(&updateTests).Error; tx != nil {
		return updateTests, tx
	}
	return updateTests, nil
}

//get 1 specified test with test struct output
func GetUpdateTests(id int) models.Tests {
	var tests models.Tests
	config.DB.Find(&tests, "id=?", id)
	return tests
}
