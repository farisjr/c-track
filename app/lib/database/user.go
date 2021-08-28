package database

import (
	"app/config"
	"app/middlewares"
	"app/models"
	"errors"
	"fmt"
)

func LoginUser(userId int, password string) (models.User, error) {
	var err error
	var user models.User
	if err = config.DB.Where("user_id=? AND password=?", userId, password).First(&user).Error; err != nil {
		return user, err
	}

	user.Token, err = middlewares.CreateToken(int(user.UserID))
	if err != nil {
		return user, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return user, err
	}
	return user, err
}

func GetOneUser(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "user_id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) (models.User, error, error) {
	userFromDb, err := GetOneUser(user.UserID)
	if err != nil {
		return userFromDb, err, nil
	}
	if userFromDb.UserID == 0 {
		if err := config.DB.Save(&user).Error; err != nil {
			fmt.Println("masuk kesini", err)
			return user, err, nil
		} else {
			switch user.Role {
			case "Patient":
				var patient models.Patient
				patient.UserID = user.UserID
				_, errCreatePatient := CreatePatient(patient)
				if errCreatePatient != nil {
					return userFromDb, nil, errCreatePatient
				}
			case "Doctor":
				var doctor models.Doctor
				doctor.UserID = user.UserID
				_, errCreateDoctor := CreateDoctor(doctor)
				if errCreateDoctor != nil {
					return userFromDb, nil, errCreateDoctor
				}
			case "Checker":
				var checker models.Checker
				checker.UserID = user.UserID
				_, errCreateChecker := CreateChecker(checker)
				if errCreateChecker != nil {
					return userFromDb, nil, errCreateChecker
				}
			}
		}
	} else {
		fmt.Println("masuk kesini2")
		return userFromDb, errors.New(""), nil
	}
	return user, nil, nil
}

func UpdateUser(users models.User, id int) (models.User, error) {
	if err := config.DB.Find(&users, "user_id=?", id).Save(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func GetUser(id int) (models.User, error) {
	var users models.User
	if err := config.DB.Find(&users, "user_id=?", id).Error; err != nil {
		return users, err
	}
	return users, nil
}
func GetDetailUser(userId int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func EditUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//get token user
func GetToken(user_id int) (string, error) {
	var user models.User
	if tx := config.DB.Model(&user).Select("token").Where("user_id=?", user_id).Error; tx != nil {
		return user.Token, tx
	}
	return user.Token, nil
}
