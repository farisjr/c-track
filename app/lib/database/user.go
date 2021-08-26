package database

import (
	"app/config"
	"app/middlewares"
	"app/models"
)

func LoginUser(username, password string) (models.User, error) {
	var err error
	var user models.User
	if err = config.DB.Where("username=? AND password=?", username, password).First(&user).Error; err != nil {
		return user, err
	}

	user.Token, err = middlewares.CreateToken(int(user.UserID))
	if err != nil {
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

func CreateUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
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
