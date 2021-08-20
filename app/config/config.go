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
	connectionString := "root:toor@tcp(localhost:3306)/c-track?charset=utf8&parseTime=True&loc=Local"
	//connectionString := os.Getenv("CONNECTION_STRING")
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
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Checker{})
	DB.AutoMigrate(&models.Patient{})
	DB.AutoMigrate(&models.Doctor{})
	DB.AutoMigrate(&models.DoctorTestDetails{})
	//DB.AutoMigrate(&models.CheckerTestDetails{})
	DB.AutoMigrate(&models.TestCategories{})
	DB.AutoMigrate(&models.Tests{})
	DB.AutoMigrate(&models.User{})
}
