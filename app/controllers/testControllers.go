package controllers

import (
	"app/lib/database"
	"app/middlewares"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func AuthorizedDoctor(c echo.Context) bool {
	_, role := middlewares.ExtractTokenUserId(c)

	if role != "Doctor" {
		return false
	}
	return true
}

//Create new test
func CreateTest(c echo.Context) error {
	auth := AuthorizedDoctor(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this page")
	}
	addTest := models.Tests{}
	c.Bind(&addTest)
	test, err := database.CreateTest(addTest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot insert data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new test",
		"data":    test,
	})
}

//get all test data
// func GetAllTests(c echo.Context) error {
// 	loggedInPatient := middlewares.ExtractTokenUserId(c)
// 	tests, err := database.GetAllTests(loggedInPatient)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	if len(tests) == 0 {
// 		return echo.NewHTTPError(http.StatusBadRequest, "No test data")
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get all test data",
// 		"data":    tests,
// 	})
// }

func GetOneTestController(c echo.Context) error {
	testId, err := strconv.Atoi(c.Param("test_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	tests, err := database.GetOneTest(testId)
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

func UpdateTest(c echo.Context) error {
	auth := AuthorizedDoctor(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}

	testId, err := strconv.Atoi(c.Param("test_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	//get test by id
	test := database.GetUpdateTest(testId)

	c.Bind(&test)
	updatedTest, err := database.UpdateTest(test)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot edit test",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":          "success update test ",
		"update test data": updatedTest,
	})
}
