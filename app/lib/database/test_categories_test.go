package database

import (
	"app/config"
	"app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBTestCategories = models.TestCategories{
		Name: "SWAB Antigen",
	}
)

func TestCreateTestCategories(t *testing.T) {
	config.InitDBTest()                                        // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.TestCategories{}) // create table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	createdTestCategories, err := CreateTestCategories(mockDBTestCategories)
	// check and test TestCategories data, if data injection exist in TestCategories's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, "SWAB Antigen", createdTestCategories.Name)
	}
}

func TestGetTestCategories(t *testing.T) {
	config.InitDBTest()                                        // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.TestCategories{}) // create table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	CreateTestCategories(mockDBTestCategories)
	// get TestCategories data from database
	getCategories, err := GetTestCategories()
	// check and test TestCategories data, if data injection exist in TestCategories's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, "SWAB Antigen", getCategories.Name)
	}
}
func TestGetTestCategoriesById(t *testing.T) {
	config.InitDBTest()                                        // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.TestCategories{}) // create table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	createdTestCategories, _ := CreateTestCategories(mockDBTestCategories)
	// get TestCategories data by id from database
	getOneTestCategories, err := GetTestCategory(int(createdTestCategories.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "SWAB Antigen", getOneTestCategories.Name)
	}
}

func TestDeleteTestCategory(t *testing.T) {
	config.InitDBTest()                                        // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.TestCategories{}) // create table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	createdTestCategories, _ := CreateTestCategories(mockDBTestCategories)
	// get TestCategories data by id from database
	testCategory, _ := GetTestCategory(int(createdTestCategories.ID))
	deletedTestCategory, err := DeleteTestCategory(int(testCategory.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "SWAB Antigen", deletedTestCategory.Name)
	}
}

func TestUpdateTestCategory(t *testing.T) {
	config.InitDBTest()                                        // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.TestCategories{}) // create table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	createdTestCategories, _ := CreateTestCategories(mockDBTestCategories)
	// get TestCategories data by id from database
	testCategory, _ := GetTestCategory(int(createdTestCategories.ID))
	// update TestCategory
	testCategory.Name = "SWAB PCR"
	// inject update TestCategories data into TestCategories's table
	updatetestCategory, err := UpdateTestCategory(testCategory)
	if assert.NoError(t, err) {
		assert.Equal(t, "SWAB PCR", updatetestCategory.Name)
	}
}
