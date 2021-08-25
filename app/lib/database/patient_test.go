package database

import (
	"app/config"
	"app/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBPatient = models.Patient{
		Dob:            time.Now(),
		Pob:            "Surabaya",
		Address:        "Surabaya",
		City:           "Surabaya",
		Province:       "Jawa Timur",
		Gender:         "Male",
		Blood_type:     "O",
		Religion:       "Kristen",
		Marital_Status: "Single",
		User:           mockDBUser,
	}
)

func TestCreatePatientSuccess(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Patient{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Patient{}) // create table from database
	// inject patient data from MockDBPatient into patient's table
	createdPatient, err := CreatePatient(mockDBPatient)
	// check and test patient data, if data injection exist in patient's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, "Surabaya", createdPatient.Pob)
		assert.Equal(t, "Surabaya", createdPatient.Address)
		assert.Equal(t, "Surabaya", createdPatient.City)
		assert.Equal(t, "Jawa Timur", createdPatient.Province)
		assert.Equal(t, models.Male, createdPatient.Gender)
		assert.Equal(t, models.O, createdPatient.Blood_type)
		assert.Equal(t, models.Kristen, createdPatient.Religion)
		assert.Equal(t, models.Single, createdPatient.Marital_Status)
	}
}

func TestCreatePatientFail(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Patient{}) // delete table from database
	// inject patient data from MockDBPatient into patient's table
	_, err := CreatePatient(mockDBPatient)
	// check and test patient data, if data injection exist in patient's table database, test will be pass
	assert.NoError(t, err)
}

func TestGetPatientSuccess(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Patient{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Patient{}) // create table from database
	// inject patient data from MockDBPatient into patient's table
	CreatePatient(mockDBPatient)
	// get patient data from database
	getPatient, err := GetPatient()
	// check and test patient data, if data exist in patient's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, "Surabaya", getPatient.Pob)
		assert.Equal(t, "Surabaya", getPatient.Address)
		assert.Equal(t, "Surabaya", getPatient.City)
		assert.Equal(t, "Jawa Timur", getPatient.Province)
		assert.Equal(t, models.Male, getPatient.Gender)
		assert.Equal(t, models.O, getPatient.Blood_type)
		assert.Equal(t, models.Kristen, getPatient.Religion)
		assert.Equal(t, models.Single, getPatient.Marital_Status)
	}
}

func TestGetPatientFail(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Patient{}) // delete table from database
	// inject patient data from MockDBPatient into patient's table
	CreatePatient(mockDBPatient)
	// get patient data from database
	_, err := GetPatient()
	// check and test patient data, if data exist in patient's table database, test will be pass
	assert.NoError(t, err)
}

func TestGetPatientByIdSuccess(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Patient{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Patient{}) // create table from database
	// inject patient data from MockDBPatient into patient's table
	createdPatient, _ := CreatePatient(mockDBPatient)
	// get patient data by id from database
	getOnePatient, err := GetPatientById(int(createdPatient.PatientID))
	// check and test patient data, if data exist in patient's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1, int(getOnePatient.PatientID))
		assert.Equal(t, "Surabaya", getOnePatient.Pob)
		assert.Equal(t, "Surabaya", getOnePatient.Address)
		assert.Equal(t, "Surabaya", getOnePatient.City)
		assert.Equal(t, "Jawa Timur", getOnePatient.Province)
		assert.Equal(t, models.Male, getOnePatient.Gender)
		assert.Equal(t, models.O, getOnePatient.Blood_type)
		assert.Equal(t, models.Kristen, getOnePatient.Religion)
		assert.Equal(t, models.Single, getOnePatient.Marital_Status)
	}
}

func TestGetPatientByIdFail(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Patient{}) // delete table from database
	// inject patient data from MockDBPatient into patient's table
	CreatePatient(mockDBPatient)
	// get patient data by id from database
	_, err := GetPatientById(1)
	// check and test patient data, if data exist in patient's table database, test will be pass
	assert.NoError(t, err)
}

func TestUpdatePatientSuccess(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Patient{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Patient{}) // create table from database
	// inject patient data from MockDBPatient into patient's table
	createdPatient, _ := CreatePatient(mockDBPatient)
	// get patient data by id from database
	patient, _ := GetPatientById(int(createdPatient.PatientID))
	// update patient data into patient's table
	patient.Address = "Bandung"
	patient.City = "Bandung"
	patient.Province = "Jawa Barat"
	// inject update patient data into patient's table
	updatePatient, err := UpdatePatient(patient)
	// check and test patient data, if data exist in patient's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1, int(updatePatient.PatientID))
		assert.Equal(t, "Surabaya", updatePatient.Pob)
		assert.Equal(t, "Bandung", updatePatient.Address)
		assert.Equal(t, "Bandung", updatePatient.City)
		assert.Equal(t, "Jawa Barat", updatePatient.Province)
		assert.Equal(t, models.Male, updatePatient.Gender)
		assert.Equal(t, models.O, updatePatient.Blood_type)
		assert.Equal(t, models.Kristen, updatePatient.Religion)
		assert.Equal(t, models.Single, updatePatient.Marital_Status)
	}
}

func TestUpdatePatientFail(t *testing.T) {
	config.InitDBTest()                               // connect to database
	config.DB.Migrator().DropTable(&models.Patient{}) // delete table from database
	_, err := UpdatePatient(mockDBPatient)
	// check and test patient data, if data exist in patient's table database, test will be pass
	assert.NoError(t, err)
}

// func TestDeletePatient(t *testing.T) {
// 	config.InitDBTest()                                 // connect to database
// 	config.DB.Migrator().DropTable(&models.Patient{})   // delete table from database
// 	config.DB.Migrator().AutoMigrate(&models.Patient{}) // create table from database
// 	// inject patient data from MockDBPatient into patient's table
// 	createdPatient, _ := CreatePatient(mockDBPatient)
// 	// get patient data by id from database
// 	patient, _ := GetPatientById(int(createdPatient.ID))
// 	// delete patient data from database, data still exist in database because of using soft delete
// 	deletedPatient, err := DeletePatient(int(patient.ID))
// 	// check and test patient data, if data still exist in patient's table database, test will be pass
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, 1, int(deletedPatient.ID))
// 		assert.Equal(t, "Surabaya", deletedPatient.Pob)
// 		assert.Equal(t, "Surabaya", deletedPatient.Address)
// 		assert.Equal(t, "Surabaya", deletedPatient.City)
// 		assert.Equal(t, "Jawa Timur", deletedPatient.Province)
// 		assert.Equal(t, models.Male, deletedPatient.Gender)
// 		assert.Equal(t, models.O, deletedPatient.Blood_type)
// 		assert.Equal(t, models.Kristen, deletedPatient.Religion)
// 		assert.Equal(t, models.Single, deletedPatient.Marital_Status)
// 	}
// }

/*
func TestPatientLogin(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Patient{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Patient{}) // create table from database
	user, _ := CreateUser(mockDBUser)
	loggedPatient, err := PatientLoginDB(user.Username, user.Password)
	if assert.NoError(t, err) {
		assert.Equal(t, "Surabaya", loggedPatient.Pob)
		assert.Equal(t, "Surabaya", loggedPatient.Address)
		assert.Equal(t, "Surabaya", loggedPatient.City)
		assert.Equal(t, "Jawa Timur", loggedPatient.Province)
		assert.Equal(t, models.Male, loggedPatient.Gender)
		assert.Equal(t, models.O, loggedPatient.Blood_type)
		assert.Equal(t, models.Kristen, loggedPatient.Religion)
		assert.Equal(t, models.Single, loggedPatient.Marital_Status)
	}
}
*/
