package database

import (
	"app/config"
	"app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBTest = models.Tests{
		Result:         "Positive",
		TestCategories: models.TestCategories{},
		Patient:        models.Patient{},
	}

	mockDBTestEdit = models.Tests{
		Result:         "Negative",
		TestCategories: models.TestCategories{},
		Patient:        models.Patient{},
	}
)

func TestCreateTestSuccess(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Tests{})
	config.DB.Migrator().AutoMigrate(&models.Tests{})
	createdTest, err := CreateTest(mockDBTest)
	if assert.NoError(t, err) {
		assert.Equal(t, "Positive", createdTest.Result)
		assert.Equal(t, "111", createdTest.TestCategories)
		assert.Equal(t, "123", createdTest.Patient)
	}
}

func TestCreateTestError(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Tests{})
	_, err := CreateTest(mockDBTest)
	assert.Error(t, err)
}

func TestGetOneTestSuccess(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Tests{})
	config.DB.Migrator().AutoMigrate(&models.Tests{})
	createdTest, _ := CreateTest(mockDBTest)
	test, err := GetOneTest(int(createdTest.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Negative", test.Result)
		assert.Equal(t, "12", test.TestCategories)
		assert.Equal(t, "123", test.Patient)
	}
}

func TestGetOneTestError(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Tests{})
	CreateTest(mockDBTest)
	_, err := GetOneTest(1)
	assert.Error(t, err)
}

func TestEditTestSuccess(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Tests{})
	config.DB.Migrator().AutoMigrate(&models.Tests{})
	createdTest, _ := CreateTest(mockDBTest)
	test, _ := GetOneTest(int(createdTest.ID))
	test.Result = "Null"
	editTest, _ := UpdateTests(test)
	testEdited, err := GetOneTest(int(editTest.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Null", testEdited.Result)
	}
}

func TestEditTestError(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Tests{})
	_, err := UpdateTests(mockDBTestEdit)
	assert.Error(t, err)
}
