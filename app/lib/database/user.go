package database

import (
	"app/config"
	"app/models"
)

func CheckSameUserId(userid int) (bool, error) {
	var user models.User
	if err := config.DB.Raw("select * from users where user_id = ?", userid).Scan(&user).Error; err != nil {
		return true, err
	}
	if user.UserID == userid {
		return true, nil
	}
	return false, nil
}

// func CreateUser(user models.User) (models.User, error, error) {
// 	userFromDb, err := GetOneUser(user.UserID)
// 	if err != nil {
// 		return userFromDb, err, nil
// 	}
// 	if userFromDb.UserID == 0 {
// 		if err := config.DB.Save(&user).Error; err != nil {
// 			fmt.Println("masuk kesini", err)
// 			return user, err, nil
// 		} else {
// 			switch user.Role {
// 			case "Patient":
// 				var patient models.Patient
// 				patient.UserID = user.UserID
// 				_, errCreatePatient := CreatePatient(patient)
// 				if errCreatePatient != nil {
// 					return userFromDb, nil, errCreatePatient
// 				}
// 			case "Doctor":
// 				var doctor models.Doctor
// 				doctor.UserID = user.UserID
// 				_, errCreateDoctor := CreateDoctor(doctor)
// 				if errCreateDoctor != nil {
// 					return userFromDb, nil, errCreateDoctor
// 				}
// 			case "Checker":
// 				var checker models.Checker
// 				checker.UserID = user.UserID
// 				_, errCreateChecker := CreateChecker(checker)
// 				if errCreateChecker != nil {
// 					return userFromDb, nil, errCreateChecker
// 				}
// 			}
// 		}
// 	} else {
// 		fmt.Println("masuk kesini2")
// 		return userFromDb, errors.New(""), nil
// 	}
// 	return user, nil, nil
// }

func UpdateUser(users models.User, id int) (models.User, error) {
	if err := config.DB.Find(&users, "user_id=?", id).Save(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

//get token user
func GetToken(user_id int) (string, error) {
	var user models.User
	if tx := config.DB.Model(&user).Select("token").Where("user_id=?", user_id).Error; tx != nil {
		return user.Token, tx
	}
	return user.Token, nil
}

func GetOneUser(userId int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "user_id=?", userId).Error; err != nil {
		return user, err
	}
	return user, nil
}
