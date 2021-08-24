package database

import (
	"app/config"
	"app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBDoctor = models.Doctor{
		Name:                   "dr. Doni Pebruwantoro",
		MedicalFacilityName:    "RS Husada Utama Surabaya",
		MedicalFacilityAddress: "Jl. Tunjungan Kota Surabaya",
		User: models.User{
			Username: "jankuat@gmail.com",
			Password: "smartdoctor",
			Role:     models.Role("Doctor"),
			Token:    "1919191919kdaieiauen",
		},
	}
)

func TestCreateDoctor(t *testing.T) {
	config.InitDBTest()                                // connect to database
	config.DB.Migrator().DropTable(&models.Doctor{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Doctor{}) // create table from database
	// inject doctor data from MockDBDoctor into doctor's table
	createdDoctor, err := CreateDoctor(mockDBDoctor)
	// check and test doctor data, if data injection exist in doctor's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, "dr. Doni Pebruwantoro", createdDoctor.Name)
		assert.Equal(t, "RS Husada Utama Surabaya", createdDoctor.MedicalFacilityName)
		assert.Equal(t, "Jl. Tunjungan Kota Surabaya", createdDoctor.MedicalFacilityAddress)
	}
}

func TestGetDoctor(t *testing.T) {
	config.InitDBTest()                                // connect to database
	config.DB.Migrator().DropTable(&models.Doctor{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Doctor{}) // create table from database
	// inject doctor data from MockDBDoctor into doctor's table
	CreateDoctor(mockDBDoctor)
	// get doctor data from database
	getDoctor, err := GetDoctor()
	// check and test doctor data, if data exist in doctor's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, "dr. Doni Pebruwantoro", getDoctor.Name)
		assert.Equal(t, "RS Husada Utama Surabaya", getDoctor.MedicalFacilityName)
		assert.Equal(t, "Jl. Tunjungan Kota Surabaya", getDoctor.MedicalFacilityAddress)
	}
}

func TestGetDoctorById(t *testing.T) {
	config.InitDBTest()                                // connect to database
	config.DB.Migrator().DropTable(&models.Doctor{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Doctor{}) // create table from database
	// inject doctor data from MockDBDoctor into doctor's table
	createdDoctor, _ := CreateDoctor(mockDBDoctor)
	// get doctor data by id from database
	getOneDoctor, err := GetDoctorById(int(createdDoctor.ID))
	// check and test doctor data, if data exist in doctor's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1, int(getOneDoctor.ID))
		assert.Equal(t, "dr. Doni Pebruwantoro", getOneDoctor.Name)
		assert.Equal(t, "RS Husada Utama Surabaya", getOneDoctor.MedicalFacilityName)
		assert.Equal(t, "Jl. Tunjungan Kota Surabaya", getOneDoctor.MedicalFacilityAddress)
	}
}

func TestUpdateDoctor(t *testing.T) {
	config.InitDBTest()                                // connect to database
	config.DB.Migrator().DropTable(&models.Doctor{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Doctor{}) // create table from database
	// inject doctor data from MockDBDoctor into doctor's table
	createdDoctor, _ := CreateDoctor(mockDBDoctor)
	// get doctor data by id from database
	doctor, _ := GetDoctorById(int(createdDoctor.ID))
	// update doctor data into checker's table
	doctor.MedicalFacilityName = "RSUD Nganjuk"
	doctor.MedicalFacilityAddress = "Jl. WR. Supratman Nganjuk"
	// inject update doctor data into doctor's table
	updateDoctor, err := UpdateDoctor(doctor)
	// check and test doctor data, if data exist in doctor's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1, int(updateDoctor.ID))
		assert.Equal(t, "dr. Doni Pebruwantoro", updateDoctor.Name)
		assert.Equal(t, "RSUD Nganjuk", updateDoctor.MedicalFacilityName)
		assert.Equal(t, "Jl. WR. Supratman Nganjuk", updateDoctor.MedicalFacilityAddress)
	}
}

func TestDeleteDoctor(t *testing.T) {
	config.InitDBTest()                                // connect to database
	config.DB.Migrator().DropTable(&models.Doctor{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Doctor{}) // create table from database
	// inject doctor data from MockDBDoctor into doctor's table
	createdDoctor, _ := CreateDoctor(mockDBDoctor)
	// get doctor data by id from database
	doctor, _ := GetDoctorById(int(createdDoctor.ID))
	// delete doctor data from database, data still exist in database because of using soft delete
	deletedDoctor, err := DeleteDoctorById(doctor)
	// check and test doctor data, if data still exist in doctor's table database, test will be pass
	if assert.NoError(t, err) {
		assert.Equal(t, 1, int(deletedDoctor.ID))
		assert.Equal(t, "dr. Doni Pebruwantoro", deletedDoctor.Name)
		assert.Equal(t, "RSUD Nganjuk", deletedDoctor.MedicalFacilityName)
		assert.Equal(t, "Jl. WR. Supratman Nganjuk", deletedDoctor.MedicalFacilityAddress)
	}
}
func TestCheckerLogin(t *testing.T) {
	config.InitDBTest()                                 // connect to database
	config.DB.Migrator().DropTable(&models.Checker{})   // delete table from database
	config.DB.Migrator().AutoMigrate(&models.Checker{}) // create table from database

}
