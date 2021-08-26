package database

import (
	"app/config"
	"app/models"
)

func CreateTest(createTests models.Tests) (models.Tests, error) {
	if err := config.DB.Save(&createTests).Error; err != nil {
		return createTests, err
	}
	return createTests, nil
}

func GetAllTests() (models.Tests, error) {
	var tests models.Tests
	if err := config.DB.Find(&tests).Error; err != nil {
		return tests, err
	}
	return tests, nil
}

func GetOneTest(id int) (models.Tests, error) {
	var tests models.Tests
	if err := config.DB.Find(&tests, "test_id=?", id).Error; err != nil {
		return tests, err
	}
	return tests, nil
}

// func DeleteTest(deleteTest models.Tests) (models.Tests, error) {
// 	if err := config.DB.Delete(&deleteTest).Error; err != nil {
// 		return deleteTest, err
// 	}
// 	return deleteTest, nil
// }

//update test info from database
func UpdateTests(updateTests models.Tests) (models.Tests, error) {
	if err := config.DB.Save(&updateTests).Error; err != nil {
		return updateTests, err
	}
	return updateTests, nil
}
