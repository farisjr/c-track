package controllers

import (
	"app/lib/database"
	"app/middlewares"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Register checker controller for checker registration
func CheckerSignUp(c echo.Context) error {
	input := models.User{}
	c.Bind(&input)
	if input.UserID == 0 || input.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please fill userid and password correctly",
		})
	}
	if same, _ := database.CheckSameId(input.UserID); same == true {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "userid already used",
		})
	}
	addChecker := models.User{}
	addChecker.UserID = input.UserID
	addChecker.Password = ourEncrypt(input.Password)
	addChecker.Role = "Checker"
	c.Bind(&addChecker)
	checker, err := database.CreateChecker(addChecker)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapChecker := map[string]interface{}{
		"ID": checker.UserID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new checker",
		"data":    mapChecker,
	})
}

//Login for checker with matching userid and password
func CheckerLogin(c echo.Context) error {
	checker := models.User{}
	c.Bind(&checker)
	loginchecker, err := database.CheckerLoginDB(checker.UserID, checker.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mapLoginchecker := map[string]interface{}{
		"User ID": loginchecker.UserID,
		"Token":   loginchecker.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes login",
		"checker": mapLoginchecker,
	})
}

//Authorization checker
func AuthorizationChecker(checkerId int, c echo.Context) error {
	authchecker, err := database.GetOneChecker(checkerId)
	loggedInCheckerId, role := middlewares.ExtractTokenUserId(c)
	if loggedInCheckerId != checkerId || string(authchecker.Role) != role || err != nil || authchecker.Role != "Checker" {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user does not have access")
	}
	return nil
}

//Logout checker
func LogoutChecker(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, _ := database.GetOneChecker(userId)
	logout.Token = ""
	checker, err := database.UpdateChecker(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Thank you",
		"data":    checker,
	})
}
