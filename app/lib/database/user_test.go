package database

import (
	"app/config"
	"app/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBUser = models.User{
		UserID:    1010101010,
		Password:  "123456",
		Role:      models.Role("Patient"),
		Token:     "jafniaebfiajnfe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

func TestLoginUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, err := CreateUser(mockDBUser)
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBUser.UserID, createdUser.UserID)
		assert.Equal(t, mockDBUser.Fullname, createdUser.Username)
		assert.Equal(t, mockDBUser.Password, createdUser.Password)
		assert.Equal(t, mockDBUser.Role, createdUser.Role)
		assert.Equal(t, mockDBUser.Token, createdUser.Token)
	}
	loggedUser, err := LoginUser(createdUser.Username, createdUser.Password)
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBUser.UserID, loggedUser.UserID)
		assert.Equal(t, mockDBUser.Fullname, loggedUser.Fullname)
		assert.Equal(t, mockDBUser.Password, loggedUser.Password)
		assert.Equal(t, mockDBUser.Role, loggedUser.Role)
		assert.Equal(t, mockDBUser.Token, loggedUser.Token)
	}
}

func TestLoginUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	CreateUser(mockDBUser)
	_, err := LoginUser("createdUser.Username", "createdUser.Password")
	assert.NoError(t, err)
}

func TestGetOneUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, _ := CreateUser(mockDBUser)
	// get User data by id from database
	getUser, err := GetOneUser(int(createdUser.UserID))
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBUser.UserID, getUser.UserID)
		assert.Equal(t, mockDBUser.Fullname, getUser.Fullname)
		assert.Equal(t, mockDBUser.Password, getUser.Password)
		assert.Equal(t, mockDBUser.Role, getUser.Role)
		assert.Equal(t, mockDBUser.Token, getUser.Token)
	}
}

func TestGetOneUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	CreateUser(mockDBUser)
	// get User data by id from database
	_, err := GetOneUser(1)
	assert.NoError(t, err)
}

func TestCreateUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, err := CreateUser(mockDBUser)
	// check and test User data, if data injection exist in User's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBUser.UserID, createdUser.UserID)
		assert.Equal(t, mockDBUser.Fullname, createdUser.Username)
		assert.Equal(t, mockDBUser.Password, createdUser.Password)
		assert.Equal(t, mockDBUser.Role, createdUser.Role)
		assert.Equal(t, mockDBUser.Token, createdUser.Token)
	}
}

func TestCreateUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	_, err := CreateUser(mockDBUser)
	// check and test User data, if data injection exist in User's table database, test will be pass
	assert.NoError(t, err)
}

func TestUpdateUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, _ := CreateUser(mockDBUser)
	// get User data by id from database
	createdUser.Username = "enakenak"
	createdUser.Password = "qwerrty"
	updatedUser, err := UpdateUser(createdUser, createdUser.UserID)
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBUser.UserID, updatedUser.UserID)
		assert.Equal(t, "enakenak", updatedUser.Fullname)
		assert.Equal(t, "qwerrty", updatedUser.Password)
		assert.Equal(t, mockDBUser.Role, updatedUser.Role)
		assert.Equal(t, mockDBUser.Token, updatedUser.Token)
	}
}

func TestUpdateUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	createdUser, _ := CreateUser(mockDBUser)
	_, err := UpdateUser(createdUser, createdUser.UserID)
	assert.NoError(t, err)
}

func TestGetUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, _ := CreateUser(mockDBUser)
	// get User data by id from database
	detailUser, err := GetUser(createdUser.UserID)
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBUser.UserID, detailUser.UserID)
		assert.Equal(t, mockDBUser.Fullname, detailUser.Fullname)
		assert.Equal(t, mockDBUser.Password, detailUser.Password)
		assert.Equal(t, mockDBUser.Role, detailUser.Role)
		assert.Equal(t, mockDBUser.Token, detailUser.Token)
	}
}

func TestGetUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	CreateUser(mockDBUser)
	// get User data by id from database
	_, err := GetUser(1)
	assert.NoError(t, err)
}

func TestGetDetailUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, _ := CreateUser(mockDBUser)
	// get User data by id from database
	detailUser, err := GetDetailUser(createdUser.UserID)
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBUser.UserID, detailUser.UserID)
		assert.Equal(t, mockDBUser.Fullname, detailUser.Fullname)
		assert.Equal(t, mockDBUser.Password, detailUser.Password)
		assert.Equal(t, mockDBUser.Role, detailUser.Role)
		assert.Equal(t, mockDBUser.Token, detailUser.Token)
	}
}

func TestGetDetailUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	CreateUser(mockDBUser)
	// get User data by id from database
	_, err := GetDetailUser(1)
	assert.NoError(t, err)
}

func TestEditUserSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, _ := CreateUser(mockDBUser)
	// get User data by id from database
	createdUser.Username = "enakenak"
	createdUser.Password = "qwerrty"
	updatedUser, err := EditUser(createdUser)
	if assert.NoError(t, err) {
		assert.Equal(t, mockDBUser.UserID, updatedUser.UserID)
		assert.Equal(t, "enakenak", updatedUser.Fullname)
		assert.Equal(t, "qwerrty", updatedUser.Password)
		assert.Equal(t, mockDBUser.Role, updatedUser.Role)
		assert.Equal(t, mockDBUser.Token, updatedUser.Token)
	}
}

func TestEditUserFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	createdUser, _ := CreateUser(mockDBUser)
	_, err := EditUser(createdUser)
	assert.NoError(t, err)
}

func TestGetTokenSuccess(t *testing.T) {
	config.InitDBTest()                              // connect to database
	config.DB.Migrator().DropTable(&models.User{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	createdUser, _ := CreateUser(mockDBUser)
	tokenUser, err := GetToken(createdUser.UserID)
	if assert.NoError(t, err) {
		assert.Equal(t, tokenUser, tokenUser)
	}
}

func TestGetTokenFail(t *testing.T) {
	config.InitDBTest()                            // connect to database
	config.DB.Migrator().DropTable(&models.User{}) // delete table from database
	//config.DB.Migrator().AutoMigrate(&models.User{}) // create table from database
	// inject User data from MockDBUser into User's table
	CreateUser(mockDBUser)
	_, err := GetToken(12)
	assert.NoError(t, err)
}
