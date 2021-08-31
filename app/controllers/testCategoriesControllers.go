package controllers

import (
	"app/lib/database"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateTestCategoriesController(c echo.Context) error {
	testCategories := models.TestCategories{}
	c.Bind(&testCategories)

	testCategoriesAdd, err := database.CreateTestCategories(testCategories)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new test categories",
		"data":    testCategoriesAdd,
	})
}

func GetAllTestCategoriesController(c echo.Context) error {
	testCategories, err := database.GetAllTestCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": testCategories,
	})
}

func GetOneTestCategoriesController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	testCategories, err := database.GetOneTestCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success get test categories by id",
		"test categories": testCategories,
	})
}

func UpdateTestCategoriesController(c echo.Context) error {
	var testCategories models.TestCategories
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	GetTestCategories, err := database.GetOneTestCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	testCategories = GetTestCategories
	c.Bind(&testCategories)
	updatedTestCategories, err := database.UpdateTestCategory(testCategories)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                "success update test categories",
		"update test categories": updatedTestCategories,
	})
}
