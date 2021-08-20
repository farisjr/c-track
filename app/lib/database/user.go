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

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return user, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetOneUser(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "userid=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(users models.User, id int) (interface{}, error) {
	if err := config.DB.Find(&users, "id=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Save(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var users models.User
	if err := config.DB.Find(&users, "id=?", id).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func GetDetailUser(userId int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}
