package controllers

import (
	"app/lib/database"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

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
