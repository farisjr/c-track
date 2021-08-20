package controllers

import (
	"app/lib/database"
	"app/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, e := database.LoginUser(user.Username, user.Password)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "username or password is not correct",
		})
	}
	mapUserLogin := map[string]interface{}{
		"ID":    user.User_ID,
		"Name":  user.Username,
		"Token": user.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "welcome",
		"user":    mapUserLogin,
	})
}

//AUTHORIZATION USER
func UserAuthorize(userId int, c echo.Context) error {
	return nil
}
