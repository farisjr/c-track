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
	c.Bind(&addUser)
	user, errUser, errInsertRelation := database.CreateUser(addUser)
	if errUser != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user already exist",
		})
	}

	if errInsertRelation != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed add data user",
		})
	}

	mapUserRegister := map[string]interface{}{
		"ID":   user.UserID,
		"Role": user.Role,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "add new user",
		"data":    mapUserRegister,
	})

}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	user, err := database.LoginUser(user.UserID, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "username or password is not correct",
		})
	}
	mapUserLogin := map[string]interface{}{
		"ID":    user.UserID,
		"Role":  user.Role,
		"Token": user.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "welcome",
		"data":    mapUserLogin,
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
		"data":    user.UserID,
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

func GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	user, err := database.GetOneUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get doctor  by id",
		"data":    user,
	})
}
