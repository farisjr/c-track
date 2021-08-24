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
			UserID:   10101010,
			Username: "gun@gmail.com",
			Password: "123456",
			Role:     models.Role("Checker"),
			Token:    "1919191919kdaieiauen",
		},
	}
)

func TestCreateCheckerSuccess(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database
	// inject checker data from MockDBChecker into checker's table
	createdChecker, err := CreateChecker(mockDBChecker)
	// check and test checker data, if data injection exist in checker's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1234567890, createdChecker.EmployeeID)
		assert.Equal(t, "Gunawan Nur Cahyo", createdChecker.Name)
		assert.Equal(t, "Mall Tunjungan Plaza", createdChecker.OfficeName)
		assert.Equal(t, "Jl. Tunjungan Kota Surabaya", createdChecker.OfficeAddress)
	}
}
func TestCreateCheckerFail(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Checker{}) // delete table from database
	_, err := CreateChecker(mockDBChecker)
	// check and test checker data, if data injection exist in checker's table database, test will be pass
	assert.NoError(t, err)
}

func TestGetCheckersSuccess(t *testing.T) {
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

func TestGetCheckersFail(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Checker{}) // delete table from database
	CreateChecker(mockDBChecker)
	// get checker data from database
	_, err := GetCheckers()
	// check and test checker data, if data exist in checker's table database, test will be pass
	assert.NoError(t, err)
}

func TestGetCheckersByIdSuccess(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database
	// inject checker data from MockDBChecker into checker's table
	createdChecker, _ := CreateChecker(mockDBChecker)
	// get checker data by id from database
	getOneChecker, err := GetCheckerById(int(createdChecker.EmployeeID))
	// check and test checker data, if data exist in checker's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1, int(getOneChecker.ID))
		assert.Equal(t, 1234567890, getOneChecker.EmployeeID)
		assert.Equal(t, "Gunawan Nur Cahyo", getOneChecker.Name)
		assert.Equal(t, "Mall Tunjungan Plaza", getOneChecker.OfficeName)
		assert.Equal(t, "Jl. Tunjungan Kota Surabaya", getOneChecker.OfficeAddress)
	}
}

func TestGetCheckersByIdFail(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Checker{}) // delete table from database
	CreateChecker(mockDBChecker)
	// get checker data by id from database
	_, err := GetCheckerById(1)
	assert.NoError(t, err)
}

func TestUpdateCheckersSuccess(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database
	// inject checker data from MockDBChecker into checker's table
	createdChecker, _ := CreateChecker(mockDBChecker)
	// get checker data by id from database
	checker, _ := GetCheckerById(int(createdChecker.EmployeeID))
	// update checker data into checker's table
	checker.OfficeName = "Mall Grand City Surabaya"
	checker.OfficeAddress = "Jl. Darmo Indah"
	// inject update checker data into checker's table
	updateChecker, err := UpdateChecker(checker)
	// check and test checker data, if data exist in checker's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, "Gunawan Nur Cahyo", updateChecker.Name)
		assert.Equal(t, "Mall Grand City Surabaya", updateChecker.OfficeName)
		assert.Equal(t, "Jl. Darmo Indah", updateChecker.OfficeAddress)
	}
}

func TestUpdateCheckersFail(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Checker{}) // delete table from database
	_, err := UpdateChecker(mockDBChecker)
	assert.NoError(t, err)
}

func TestCheckersLogin(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database

}
