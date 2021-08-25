package database

import (
	"app/config"
	"app/middlewares"
	"app/models"
)

func LoginUser(username, password string) (interface{}, error) {
	var err error
	var user models.User
	if err = config.DB.Where("username=? AND password=?", username, password).First(&user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.UserID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, err
}

func GetOneUser(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "userid=?", id).Error; err != nil {
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
	if err := config.DB.Find(&users, "id=?", id).Error; err != nil {
		return users, err
	}
	if err := config.DB.Save(&users).Error; err != nil {
		return users, err
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

func EditUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//get token user
func GetToken(user_id int) string {
	var user models.User
	config.DB.Model(&user).Select("token").Where("user_id=?", user_id)
	return user.Token
}
