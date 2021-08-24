package database

import (
	"app/config"
	"app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBUser = models.User{
		UserID:   1010101010,
		Username: "boruto@gmail.com",
		Password: "123456",
		Role:     models.Role("Patient"),
		Token:    "jafniaebfiajnfe",
	}
	mockDBUserEdit = models.User{
		UserID:   1010101010,
		Username: "boruto@gmail.com",
		Password: "123456",
		Role:     models.Role("Patient"),
		Token:    "jafniaebfiajnfe",
	}
)

func TestCreateUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, err := CreateUser(mockDBUser)
	// check and test User data, if data injection exist in User's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1010101010, createdUser.UserID)
		assert.Equal(t, "boruto@gmail.com", createdUser.Username)
		assert.Equal(t, "123456", createdUser.Password)
		assert.Equal(t, models.Role("Patient"), createdUser.Role)
		assert.Equal(t, "jafniaebfiajnfe", createdUser.Token)
	}
}

func TestCreateUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	// inject User data from MockDBUser into User's table
	_, err := CreateUser(mockDBUser)
	// check and test User data, if data injection exist in User's table database, test will be pass
	assert.NoError(t, err)
}

func TestGetOneUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, _ := CreateUser(mockDBUser)
	// get User data by id from database
	getUser, err := GetOneUser(int(createdUser.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, 1010101010, getUser.UserID)
		assert.Equal(t, "boruto@gmail.com", getUser.Username)
		assert.Equal(t, "123456", getUser.Password)
		assert.Equal(t, models.Role("Patient"), getUser.Role)
		assert.Equal(t, "jafniaebfiajnfe", getUser.Token)
	}
}

func TestGetOneUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	CreateUser(mockDBUser)
	// get User data by id from database
	_, err := GetOneUser(1)
	assert.NoError(t, err)
}

func TestUpdateUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, _ := CreateUser(mockDBUser)
	// get User data by id from database
	getUser, _ := GetOneUser(int(createdUser.ID))
	getUser.Username = "enakenak"
	getUser.Password = "qwerrty"
	updatedUser, err := UpdateUser(getUser, getUser.UserID)
	if assert.NoError(t, err) {
		assert.Equal(t, 1010101010, updatedUser.UserID)
		assert.Equal(t, "enakenak", updatedUser.Username)
		assert.Equal(t, "qwerrty", updatedUser.Password)
		assert.Equal(t, models.Role("Patient"), updatedUser.Role)
		assert.Equal(t, "jafniaebfiajnfe", updatedUser.Token)
	}
}

func TestUpdateUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	_, err := EditUser(mockDBUserEdit)
	assert.NoError(t, err)
}
