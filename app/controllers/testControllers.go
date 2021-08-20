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
	tests, err := database.GetAllTests()
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
	tests, err := database.GetTestsId(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get test  by id",
		"data":    tests,
	})
}

func DeleteTestsByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	tests, err := database.GetTestsId(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	testsDeleted, err := database.DeleteTestsById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success delete selected test",
		"test before delete": tests,
		"test after delete":  testsDeleted,
	})
}

func UpdateTestsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	tests := database.GetUpdateTests(id)
	c.Bind(&tests)
	testsUpdateCategories, err1 := database.UpdateTests(tests)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt":    testsUpdateCategories.CreatedAt,
		"UpdatedAt":    testsUpdateCategories.UpdatedAt,
		"ResultStatus": tests.Result_Status,
		"DeletedAt":    testsUpdateCategories.DeletedAt,
		"id":           testsUpdateCategories.ID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":          "success update test ",
		"update test data": output,
	})
}
