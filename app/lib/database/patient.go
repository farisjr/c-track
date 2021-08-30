package database

import (
	"app/config"
	"app/middlewares"
	"app/models"
)

func CheckSameId(userId int) (bool, error) {
	var user models.User
	if err := config.DB.Raw("select * from users where user_id = ?", userId).Scan(&user).Error; err != nil {
		return true, err
	}
	if user.UserID == userId {
		return true, nil
	}
	return false, nil
}

func CreatePatient(patient models.User) (models.User, error) {
	if err := config.DB.Save(&patient).Error; err != nil {
		return patient, err
	}
	return patient, nil
}

// func GetPatient() (models.Patient, error) {
// 	var patients models.Patient
// 	if err := config.DB.Find(&patients).Error; err != nil {
// 		return patients, err
// 	}
// 	return patients, nil
// }

func GetOnePatient(id int) (models.User, error) {
	var patient models.User

	if err := config.DB.Find(&patient, "user_id=?", id).Error; err != nil {
		return patient, err
	}
	return patient, nil
}

// func UpdatePatient(patient models.Patient) (models.Patient, error) {
// 	// db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// 	if err := config.DB.Save(&patient).Error; err != nil {
// 		return patient, err
// 	}
// 	return patient, nil
// }

//Login for patient with matching userid and password
func PatientLoginDB(userId int, password string) (models.User, error) {
	var patient models.User
	var err error
	if err = config.DB.Where("user_id=? AND password=?", userId, password).First(&patient).Error; err != nil {
		return patient, err
	}
	patient.Token, err = middlewares.CreatePatientToken(int(patient.UserID))
	if err != nil {
		return patient, err
	}
	if err := config.DB.Save(patient).Error; err != nil {
		return patient, err
	}
	return patient, err
}

func UpdatePatient(patient models.User) (models.User, error) {
	if err := config.DB.Save(&patient).Error; err != nil {
		return patient, err
	}
	return patient, nil
}
