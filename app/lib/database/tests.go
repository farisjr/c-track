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

func GetAllTests() ([]models.Tests, error) {
	var tests []models.Tests
	if err := config.DB.Find(&tests).Error; err != nil {
		return nil, err
	}
	return tests, nil
}

func GetOneTest(testId int) (models.Tests, error) {
	var test models.Tests
	if err := config.DB.Where("test_id=?", testId).First(&test).Error; err != nil {
		return test, err
	}
	return test, nil
}

func GetTestbyPatient(patientId int) ([]models.Tests, error) {
	var test []models.Tests
	if err := config.DB.Where("patient_id=?", patientId).Find(&test).Error; err != nil {
		return test, err
	}
	return test, nil
}

func GetTestbyDoctor(doctorId int) ([]models.Tests, error) {
	var test []models.Tests
	if err := config.DB.Where("doctor_id=?", doctorId).Find(&test).Error; err != nil {
		return test, err
	}
	return test, nil
}

//update test info from database
func UpdateTest(test models.Tests) (interface{}, error) {
	if err := config.DB.Save(&test).Error; err != nil {
		return nil, err
	}
	return test, nil
}

func GetUpdateTest(testId int) models.Tests {
	var test models.Tests
	config.DB.Find(&test, "test_id=?", testId)
	return test
}
