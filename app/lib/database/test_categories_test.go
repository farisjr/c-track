package database

import (
	"app/config"
	"app/models"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBTestCategories = models.TestCategories{
		TestCategoriesID: 1,
		Name:             "SWAB Antigen",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		DeletedAt:        sql.NullTime{},
	}
)

func TestCreateTestCategoriesSuccess(t *testing.T) {
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

func TestCreateTestCategoriesFail(t *testing.T) {
	config.InitDBTest()                                      // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{}) // delete table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	_, err := CreateTestCategories(mockDBTestCategories)
	// check and test TestCategories data, if data injection exist in TestCategories's table database, test will be pass
	assert.NoError(t, err)
}

func TestGetTestCategoriesSuccess(t *testing.T) {
	config.InitDBTest()                                        // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.TestCategories{}) // create table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	CreateTestCategories(mockDBTestCategories)
	// get TestCategories data from database
	getCategories, err := GetAllTestCategories()
	// check and test TestCategories data, if data injection exist in TestCategories's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, "SWAB Antigen", getCategories)
	}
}

func TestGetTestCategoriesFail(t *testing.T) {
	config.InitDBTest()                                      // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{}) // delete table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	CreateTestCategories(mockDBTestCategories)
	// get TestCategories data from database
	_, err := GetAllTestCategories()
	// check and test TestCategories data, if data injection exist in TestCategories's table database, test will be pass
	assert.NoError(t, err)
}

func TestGetTestCategoriesByIdSuccess(t *testing.T) {
	config.InitDBTest()                                        // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.TestCategories{}) // create table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	CreateTestCategories(mockDBTestCategories)
	// get TestCategories data by id from database
	testCategory, err := GetOneTestCategory(1)
	if assert.NoError(t, err) {
		assert.Equal(t, "SWAB Antigen", testCategory.Name)
	}
}

func TestGetTestCategoriesByIdFail(t *testing.T) {
	config.InitDBTest()                                      // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{}) // delete table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	CreateTestCategories(mockDBTestCategories)
	// get TestCategories data by id from database
	_, err := GetOneTestCategory(1)
	assert.NoError(t, err)
}

// func TestDeleteTestCategory(t *testing.T) {
// 	config.InitDBTest()                                        // connect to database
// 	config.DB.Migrator().DropTable(&models.TestCategories{})   // delete table from database
// 	config.DB.Migrator().AutoMigrate(&models.TestCategories{}) // create table from database
// 	// inject TestCategories data from MockDBTestCategories into TestCategories's table
// 	createdTestCategories, _ := CreateTestCategories(mockDBTestCategories)
// 	// get TestCategories data by id from database
// 	testCategory, _ := GetTestCategory(int(createdTestCategories.ID))
// 	deletedTestCategory, err := DeleteTestCategory(int(testCategory.ID))
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, "SWAB Antigen", deletedTestCategory.Name)
// 	}
// }

func TestUpdateTestCategorySuccess(t *testing.T) {
	config.InitDBTest()                                        // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.TestCategories{}) // create table from database
	// inject TestCategories data from MockDBTestCategories into TestCategories's table
	createdTestCategories, _ := CreateTestCategories(mockDBTestCategories)
	// get TestCategories data by id from database
	testCategory, _ := GetOneTestCategory(int(createdTestCategories.TestCategoriesID))
	// update TestCategory
	testCategory.Name = "SWAB PCR"
	// inject update TestCategories data into TestCategories's table
	updatetestCategory, err := UpdateTestCategory(testCategory)
	if assert.NoError(t, err) {
		assert.Equal(t, "SWAB PCR", updatetestCategory.Name)
	}
}

func TestUpdateTestCategoryFail(t *testing.T) {
	config.InitDBTest()                                      // connect to database
	config.DB.Migrator().DropTable(&models.TestCategories{}) // delete table from database
	// inject update TestCategories data into TestCategories's table
	_, err := UpdateTestCategory(mockDBTestCategories)
	assert.NoError(t, err)
}
