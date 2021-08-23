package database

import (
	"app/config"
	"app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBPatient = models.Patient{
		//Dob:            2010 - 10 - 10,
		Pob:            "Surabaya",
		Address:        "Surabaya",
		City:           "Surabaya",
		Province:       "Jawa Timur",
		Gender:         "Male",
		Blood_type:     "O",
		Religion:       "Kristen",
		Marital_Status: "Mariage",
		User:           models.User{},
	}

	mockDBPatientEdit = models.Patient{
		//Dob:            2010 - 10 - 10,
		Pob:            "Surabaya",
		Address:        "Jakarta",
		City:           "Jakarta",
		Province:       "Jakarta",
		Gender:         "Male",
		Blood_type:     "O",
		Religion:       "Kristen",
		Marital_Status: "Mariage",
		User:           models.User{},
	}
)

func TestCreatePatientSuccess(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
	config.DB.Migrator().AutoMigrate(&models.Patient{})
	createdPatient, err := CreatePatientTest(mockDBPatient)
	if assert.NoError(t, err) {
		assert.Equal(t, "Surabaya", createdPatient.Pob)
		assert.Equal(t, "Surabaya", createdPatient.Address)
		assert.Equal(t, "Surabaya", createdPatient.City)
		assert.Equal(t, "Jawa Timur", createdPatient.Province)
		assert.Equal(t, "Male", createdPatient.Gender)
		assert.Equal(t, "O", createdPatient.Blood_type)
		assert.Equal(t, "Kristen", createdPatient.Religion)
		assert.Equal(t, "Mariage", createdPatient.Marital_Status)
	}
}

func TestCreatePatientError(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
	_, err := CreatePatientTest(mockDBPatient)
	assert.Error(t, err)
}

func TestGetPatientSuccess(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
	config.DB.Migrator().AutoMigrate(&models.Patient{})
	createdPatient, err := CreatePatientTest(mockDBPatient)
	if assert.NoError(t, err) {
		assert.Equal(t, "1", createdPatient.ID)
		assert.Equal(t, "2020-01-01", createdPatient.Dob)
		assert.Equal(t, "Surabaya", createdPatient.Pob)
		assert.Equal(t, "Surabaya", createdPatient.Address)
		assert.Equal(t, "Surabaya", createdPatient.City)
		assert.Equal(t, "Jawa Timur", createdPatient.Province)
		assert.Equal(t, "Male", createdPatient.Gender)
		assert.Equal(t, "O", createdPatient.Blood_type)
		assert.Equal(t, "Kristen", createdPatient.Religion)
		assert.Equal(t, "Mariage", createdPatient.Marital_Status)
	}
	getPatient, err := GetPatientTest()
	if assert.NoError(t, err) {
		assert.Equal(t, "1", createdPatient.ID)
		assert.Equal(t, "2020-01-01", createdPatient.Dob)
		assert.Equal(t, "Surabaya", getPatient.Pob)
		assert.Equal(t, "Surabaya", getPatient.Address)
		assert.Equal(t, "Surabaya", getPatient.City)
		assert.Equal(t, "Jawa Timur", getPatient.Province)
		assert.Equal(t, "Male", getPatient.Gender)
		assert.Equal(t, "O", getPatient.Blood_type)
		assert.Equal(t, "Kristen", getPatient.Religion)
		assert.Equal(t, "Mariage", getPatient.Marital_Status)
	}
}

func TestGetPatientError(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
	CreatePatient(mockDBPatient)
	_, err := GetPatient()
	assert.Error(t, err)
}

/*func TestGetPatientByIdSuccess(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
	config.DB.Migrator().AutoMigrate(&models.Patient{})
}

func TestGetPatientByIdError(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
}*/

func TestUpdatePatientSuccess(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
	config.DB.Migrator().AutoMigrate(&models.Patient{})
	createdPatient, _ := CreatePatientTest(mockDBPatient)
	patient, _ := GetPatientByIdTest(int(createdPatient.ID))
	patient.Pob = "Surabaya"
	updatePatient, _ := UpdatePatientTest(patient)
	updatedPatient, err := GetPatientByIdTest(int(updatePatient.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Surabaya", updatedPatient.Pob)
	}
}

func TestUpdatePatientError(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
	_, err := UpdatePatientTest(mockDBPatientEdit)
	assert.Error(t, err)
}

/*
func TestDeletePatientSuccess(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
	config.DB.Migrator().AutoMigrate(&models.Patient{})
}

func TestDeletePatientError(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
}

func TestPatientLoginSuccess(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
	config.DB.Migrator().AutoMigrate(&models.Patient{})
}

func TestPatientLoginError(t *testing.T) {
	config.InitDBTest()
	config.DB.Migrator().DropTable(&models.Patient{})
}
*/
