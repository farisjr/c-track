package database

import (
	"app/config"
	"app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBChecker = models.Checker{
		EmployeeID:    1234567890,
		Name:          "Gunawan Nur Cahyo",
		OfficeName:    "Mall Tunjungan Plaza",
		OfficeAddress: "Jl. Tunjungan Kota Surabaya",
		User: models.User{
			Username: "gun@gmail.com",
			Password: "123456",
			Role:     models.Role("Checker"),
			Token:    "1919191919kdaieiauen",
		},
	}
)

func TestCreateChecker(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database
	// inject checker data from MockDBChecker into checker's table
	createdChecker, err := CreateChecker(mockDBChecker)
	// check and test patient data, if data injection exist in patient's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1234567890, createdChecker.EmployeeID)
		assert.Equal(t, "Gunawan Nur Cahyo", createdChecker.Name)
		assert.Equal(t, "Mall Tunjungan Plaza", createdChecker.OfficeName)
		assert.Equal(t, "Jl. Tunjungan Kota Surabaya", createdChecker.OfficeAddress)
	}
}

func TestGetChecker(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database
	// inject checker data from MockDBChecker into checker's table
	CreateChecker(mockDBChecker)
	// get checker data from database
	getChecker, err := GetCheckers()
	// check and test checker data, if data exist in checker's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1234567890, getChecker.EmployeeID)
		assert.Equal(t, "Gunawan Nur Cahyo", getChecker.Name)
		assert.Equal(t, "Mall Tunjungan Plaza", getChecker.OfficeName)
		assert.Equal(t, "Jl. Tunjungan Kota Surabaya", getChecker.OfficeAddress)
	}
}

func TestGetCheckerById(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database
	// inject checker data from MockDBChecker into checker's table
	createdChecker, _ := CreateChecker(mockDBChecker)
	// get checker data by id from database
	getOneChecker, err := GetCheckerById(int(createdChecker.ID))
	// check and test checker data, if data exist in checker's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1, int(getOneChecker.ID))
		assert.Equal(t, 1234567890, getOneChecker.EmployeeID)
		assert.Equal(t, "Gunawan Nur Cahyo", getOneChecker.Name)
		assert.Equal(t, "Mall Tunjungan Plaza", getOneChecker.OfficeName)
		assert.Equal(t, "Jl. Tunjungan Kota Surabaya", getOneChecker.OfficeAddress)
	}
}

func TestUpdateChecker(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database
	// inject checker data from MockDBChecker into checker's table
	createdChecker, _ := CreateChecker(mockDBChecker)
	// get checker data by id from database
	checker, _ := GetCheckerById(int(createdChecker.ID))
	// update checker data into checker's table
	checker.EmployeeID = 678910111213
	checker.OfficeName = "Mall Grand City Surabaya"
	checker.OfficeAddress = "Jl. Darmo Indah"
	// inject update checker data into checker's table
	updateChecker, err := UpdateChecker(checker)
	// check and test checker data, if data exist in checker's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1, int(updateChecker.ID))
		assert.Equal(t, 678910111213, updateChecker.EmployeeID)
		assert.Equal(t, "Gunawan Nur Cahyo", updateChecker.Name)
		assert.Equal(t, "Mall Grand City Surabaya", updateChecker.OfficeName)
		assert.Equal(t, "Jl. Darmo Indah", updateChecker.OfficeAddress)
	}
}

func TestCheckerLogin(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database

}
