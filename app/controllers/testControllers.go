package controllers

import (
	"app/lib/database"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Create new test from registered patients
func CreateTestsController(c echo.Context) error {
	tests := models.Tests{}
	c.Bind(&tests)
	testsAdd, err := database.CreateTest(tests)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new test",
		"data":    testsAdd,
	})
}

//get all test data
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

func GetOneTestController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	tests, err := database.GetOneTest(id)
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

// func DeleteTestsController(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}
// 	tests, _ := database.GetOneTest(id)
// 	c.Bind(&tests)
// 	testsDeleted, err := database.DeleteTest(tests)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "cannot delete data",
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success delete selected test",
// 		"data":    testsDeleted,
// 	})
// }

func UpdateTestsController(c echo.Context) error {
	var test models.Tests
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	//get test by id
	updateTests, _ := database.GetOneTest(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	test = updateTests
	c.Bind(&test)
	updatedTest, err := database.UpdateTests(test)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":          "success update test ",
		"update test data": updatedTest,
	})
}
