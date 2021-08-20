package controllers

import (
	"app/lib/database"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateTestsController(c echo.Context) error {
	tests := models.Tests{}
	c.Bind(&tests)

	testsAdd, err := database.CreateTests(tests)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new test",
		"data":    testsAdd,
	})
}

func GetTestsController(c echo.Context) error {
	tests, err := database.GetTests()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all test data",
		"data":    tests,
	})
}

func GetTestsIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	tests, err := database.GetTestsById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get test  by id",
		"data":    tests,
	})
}

func DeleteTestsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	tests, _ := database.GetTestsById(id)
	c.Bind(&tests)
	testsDeleted, err := database.DeleteTestsById(tests)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot delete data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete selected test",
		"data":    testsDeleted,
	})
}

func UpdateTestsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	updateTests := database.GetUpdateTests(id)
	c.Bind(&updateTests)
	testsId, err1 := database.UpdateTests(updateTests)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":          "success update test ",
		"update test data": testsId,
	})
}
