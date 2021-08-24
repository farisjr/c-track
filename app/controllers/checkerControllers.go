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
func RegisterCheckerController(c echo.Context) error {
	checker := models.User{}
	checker.Role = "checker"
	c.Bind(&checker)
	addchecker, err := database.CreateChecker(models.Checker{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new checker",
		"data":    addchecker,
	})
}

//Login for checker with matching username and password
func LoginChecker(c echo.Context) error {
	checker := models.User{}
	c.Bind(&checker)
	loginchecker, err := database.LoginCheckerDB(checker.Username, checker.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mapLoginchecker := map[string]interface{}{
		"Employee ID": loginchecker.EmployeeID,
		"Name":        loginchecker.Name,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes login",
		"checker": mapLoginchecker,
	})
}

//Authorization checker
func AuthorizationChecker(checkerId int, c echo.Context) error {
	authchecker, err := database.GetCheckerById(checkerId)
	LoggedInCheckerId, role := middlewares.ExtractTokenUserId(c)
	if LoggedInCheckerId != checkerId || string(authchecker.User.Role) != role || err != nil || authchecker.User.Role != "checker" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access")
	}
	return nil
}

//Logout checker
func LogoutChecker(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("checkerId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, err := database.GetCheckerById(id)
	checker, err := database.UpdateChecker(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Thank you",
		"data":    checker.Name,
	})
}

func CreateCheckersController(c echo.Context) error {
	checkers := models.Checker{}
	c.Bind(&checkers)

	checkersAdd, err := database.CreateChecker(checkers)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success add new checker",
		"data":    checkersAdd,
	})
}

func GetCheckersController(c echo.Context) error {
	doctors, err := database.GetCheckers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get doctors data",
		"data":    doctors,
	})
}

func GetCheckerController(c echo.Context) error {
	checker, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get checker data",
		"data":    checker,
	})
}
