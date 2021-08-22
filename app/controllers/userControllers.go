package controllers

import (
	"app/lib/database"
	"app/middlewares"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func RegisterUserController(c echo.Context) error {
	addUser := models.User{}
	addUser.Role = "user"
	c.Bind(&addUser)
	user, err := database.CreateUser(addUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "add new user",
		"data":    user,
	})

}

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
		"ID":    user.UserID,
		"Name":  user.Username,
		"Token": user.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "welcome",
		"user":    mapUserLogin,
	})
}

//Authorization user
func UserAuthorize(userId int, c echo.Context) error {
	userAuth, err := database.GetOneUser(userId)
	loggedInUserId, role := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != userId || string(userAuth.Role) != role || err != nil || userAuth.Role != "patient" || userAuth.Role != "doctor" || userAuth.Role != "checker" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access")
	}
	return nil
}

//Logout user
func LogoutUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	if err = UserAuthorize(id, c); err != nil {
		return err
	}
	logout, err := database.GetOneUser(id)
	logout.Token = ""
	user, err := database.EditUser(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Thank you",
		"data":    user.Username,
	})
}

func EditUserProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if err = UserAuthorize(id, c); err != nil {
		return err
	}

	editUser, err := database.GetOneUser(id)
	c.Bind(&editUser)
	user, err := database.EditUser(editUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Profile Updated!",
		"data":    user,
	})
}
