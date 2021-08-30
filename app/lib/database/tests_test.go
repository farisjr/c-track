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
		TestsID:          12,
		PatientID:        1010101010,
		TestCategoriesID: 1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		TestCategories: models.TestCategories{
			TestCategoriesID: 1,
			Name:             "SWAB Antigen",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		Patient: models.Patient{
			PatientID:      1010101010,
			UserID:         12345,
			Dob:            time.Now(),
			Pob:            "Surabaya",
			Address:        "Surabaya",
			City:           "Surabaya",
			Province:       "Jawa Timur",
			Gender:         models.Gender("Male"),
			Blood_type:     models.Blood_type("AB"),
			Religion:       models.Religion("Kristen"),
			Marital_Status: models.Marital_Status("Single"),
			User: models.User{
				UserID:   1010101010,
				Fullname: "boruto@gmail.com",
				Password: "123456",
				Role:     models.Role("Patient"),
				Token:    "jafniaebfiajnfe",
			},
		},
		Result: "Negative",
	}
)

func TestCreateTestSuccess(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	createdTest, err := CreateTest(mockDBTest)
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBTest.TestsID, createdTest.TestsID)
		assert.Equal(t, mockDBTest.TestCategoriesID, createdTest.TestCategoriesID)
		assert.Equal(t, mockDBTest.PatientID, createdTest.PatientID)
	}
}

func TestCreateTestFail(t *testing.T) {
	config.InitDBTest()                             // connect to database
	config.DB.Migrator().DropTable(&models.Tests{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	_, err := CreateTest(mockDBTest)
	assert.NoError(t, err)
}

func TestGetAllTestSuccess(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	CreateTest(mockDBTest)
	getTest, err := GetAllTests()
	if assert.NoError(t, err) {
		assert.Equal(t, "Positive", getTest)
		assert.Equal(t, "SWAB Antigen", getTest)
		assert.Equal(t, "Surabaya", getTest)
	}
}

func TestGetAllTestFail(t *testing.T) {
	config.InitDBTest()                             // connect to database
	config.DB.Migrator().DropTable(&models.Tests{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	CreateTest(mockDBTest)
	_, err := GetAllTests()
	assert.NoError(t, err)
}

func TestGetOneTestSuccess(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	createdTest, _ := CreateTest(mockDBTest)
	// get Test data by id from database
	test, err := GetOneTest(createdTest.TestsID)
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBTest.TestCategoriesID, test.TestsID)
		assert.Equal(t, mockDBTest.TestCategoriesID, test.TestCategoriesID)
		assert.Equal(t, mockDBTest.PatientID, test.PatientID)
	}
}

func TestGetOneTestFail(t *testing.T) {
	config.InitDBTest()                             // connect to database
	config.DB.Migrator().DropTable(&models.Tests{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	CreateTest(mockDBTest)
	// get Test data by id from database
	_, err := GetOneTest(1)
	assert.NoError(t, err)
}
func TestUpdateTestSuccess(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	// inject Test data from MockDBTest into Test's table
	createdTest, _ := CreateTest(mockDBTest)
	// get Test data by id from database
	createdTest.Result = "Positive"
	// inject update Test data into Test's table
	editTest, err := UpdateTests(createdTest)
	if assert.NoError(t, err) {
		assert.Equal(t, "Positive", editTest.Result)
	}
}

func TestUpdateTestFail(t *testing.T) {
	config.InitDBTest()                             // connect to database
	config.DB.Migrator().DropTable(&models.Tests{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
	_, err := UpdateTests(mockDBTest)
	assert.NoError(t, err)
}

// func TestDeleteTestSuccess(t *testing.T) {
// 	config.InitDBTest()                               // connect to database
// 	config.DB.Migrator().DropTable(&models.Tests{})   // delete table from database
// 	config.DB.Migrator().AutoMigrate(&models.Tests{}) // create table from database
// 	// inject Test data from MockDBTest into Test's table
// 	createdTest, _ := CreateTest(mockDBTest)
// 	// get Test data by id from database
// 	test, _ := GetOneTest(int(createdTest.ID))
// 	deletedTest, err := DeleteTest(test)
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, "Positive", deletedTest.Result)
// 		assert.Equal(t, "SWAB Antigen", deletedTest.TestCategories.Name)
// 		assert.Equal(t, "Surabaya", deletedTest.Patient.City)
// 	}
// }

// func TestDeleteTestFail(t *testing.T) {
// 	config.InitDBTest()                             // connect to database
// 	config.DB.Migrator().DropTable(&models.Tests{}) // delete table from database
// 	// inject Test data from MockDBTest into Test's table
// 	CreateTest(mockDBTest)
// 	// get Test data by id from database
// 	_, err := DeleteTest(1)
// 	assert.NoError(t, err)
// }
