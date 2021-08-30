package database

import (
	"app/config"
	"app/middlewares"
	"app/models"
)

func CreateChecker(checker models.User) (models.User, error) {
	if err := config.DB.Save(&checker).Error; err != nil {
		return checker, err
	}
	return checker, nil
}

// func GetCheckers() (models.Checker, error) {
// 	var checkers models.Checker
// 	if err := config.DB.Find(&checkers).Error; err != nil {
// 		return checkers, err
// 	}
// 	return checkers, nil
// }

func GetOneChecker(id int) (models.User, error) {
	var checker models.User
	if err := config.DB.Find(&checker, "user_id=?", id).Error; err != nil {
		return checker, err
	}
	return checker, nil
}

func UpdateChecker(checker models.User) (models.User, error) {
	if err := config.DB.Save(&checker).Error; err != nil {
		return checker, err
	}
	return checker, nil
}

//Login for checker with matching email and password
func CheckerLoginDB(userid int, password string) (models.User, error) {
	var checker models.User
	var err error
	if err = config.DB.Where("user_id=? AND password=?", userid, password).First(&checker).Error; err != nil {
		return checker, err
	}
	checker.Token, err = middlewares.CreateCheckerToken(checker.UserID)
	if err != nil {
		return checker, err
	}
	if err := config.DB.Save(checker).Error; err != nil {
		return checker, err
	}
	return checker, err
}
