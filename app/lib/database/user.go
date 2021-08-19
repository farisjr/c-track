package database

import (
	"app/config"
	"app/models"
)

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
