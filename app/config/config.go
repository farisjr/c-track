package config

import (
	"app/models"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {
<<<<<<< HEAD
	//connectionString := "root:02021996Doni*@tcp(localhost:3306)/baru?charset=utf8&parseTime=True&loc=Local"
=======
>>>>>>> add-unit-test2
	connectionString := os.Getenv("CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}
func InitDBTest() {
	connectionString := os.Getenv("CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error
	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
	// HTTP_PORT = 80
=======
>>>>>>> add-unit-test2
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Checker{})
	DB.AutoMigrate(&models.Patient{})
	DB.AutoMigrate(&models.Doctor{})
	//DB.AutoMigrate(&models.DoctorTestDetails{})
	//DB.AutoMigrate(&models.CheckerTestDetails{})
	// DB.AutoMigrate(&models.TestCategories{})
	DB.AutoMigrate(&models.Tests{})
}
