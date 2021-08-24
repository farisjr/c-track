package database

import (
	"app/config"
	"app/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBTest = models.Tests{
		Result: "Positive",
		TestCategories: models.TestCategories{
			Name: "SWAB Antigen",
		},
		Patient: models.Patient{
			Dob:            time.Now(),
			Pob:            "Surabaya",
			Address:        "Surabaya",
			City:           "Surabaya",
			Province:       "Jawa Timur",
			Gender:         "Male",
			Blood_type:     "O",
			Religion:       "Kristen",
			Marital_Status: "Single",
			User: models.User{
				UserID:   1010101010,
				Username: "boruto@gmail.com",
				Password: "123456",
				Role:     models.Role("Patient"),
				Token:    "jafniaebfiajnfe",
			},
		},
	}
)

func TestCreateTest(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	createdTest, err := CreateTest(mockDBTest)
	if assert.NoError(t, err) {
		assert.Equal(t, "Positive", createdTest.Result)
		assert.Equal(t, "SWAB Antigen", createdTest.TestCategories.Name)
		assert.Equal(t, "Surabaya", createdTest.Patient.City)
	}
}

func TestGetAllTest(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	CreateTest(mockDBTest)
	getTest, err := GetAllTests()
	if assert.NoError(t, err) {
		assert.Equal(t, "Positive", getTest.Result)
		assert.Equal(t, "SWAB Antigen", getTest.TestCategories.Name)
		assert.Equal(t, "Surabaya", getTest.Patient.City)
	}
}

func TestGetOneTest(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	createdTest, _ := CreateTest(mockDBTest)
	// get Test data by id from database
	test, err := GetOneTest(int(createdTest.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Positive", test.Result)
		assert.Equal(t, "SWAB Antigen", test.TestCategories.Name)
		assert.Equal(t, "Surabaya", test.Patient.City)
	}
}

func TestUpdateTest(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	createdTest, _ := CreateTest(mockDBTest)
	// get Test data by id from database
	test, _ := GetOneTest(int(createdTest.ID))
	// update Test
	test.Result = "Negative"
	// inject update Test data into Test's table
	editTest, _ := UpdateTests(test)
	testEdited, err := GetOneTest(int(editTest.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Negative", testEdited.Result)
	}
}

func TestDeleteTest(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	createdTest, _ := CreateTest(mockDBTest)
	// get Test data by id from database
	test, _ := GetOneTest(int(createdTest.ID))
	deletedTest, err := DeleteTest(test)
	if assert.NoError(t, err) {
		assert.Equal(t, "Positive", deletedTest.Result)
		assert.Equal(t, "SWAB Antigen", deletedTest.TestCategories.Name)
		assert.Equal(t, "Surabaya", deletedTest.Patient.City)
	}
}
