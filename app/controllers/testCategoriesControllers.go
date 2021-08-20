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

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt":            testCategoriesAdd.CreatedAt,
		"UpdatedAt":            testCategoriesAdd.UpdatedAt,
		"DeletedAt":            testCategoriesAdd.DeletedAt,
		"id":                   testCategoriesAdd.ID,
		"test_categories_name": testCategoriesAdd.Test_categories_Name,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new test categories",
		"data":    output,
	})
}

func GetTestCategoriesController(c echo.Context) error {
	testCategories, err := database.GetAllTestCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success get all test categories",
		"test categories": testCategories,
	})
}

func GetTestCategoriesIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	testCategories, err := database.GetTestCategoriesId(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success get test categories by id",
		"test categories": testCategories,
	})
}

func DeleteTestCategoriesByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	testCategories, err := database.GetTestCategoriesId(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	testCategoriesDeleted, err := database.DeleteTestCategoriesById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                       "success delete selected test categories",
		"test categories before delete": testCategories,
		"test categories after delete":  testCategoriesDeleted,
	})
}

func UpdateTestCategoriesController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	testCategories := database.GetUpdateTestCategories(id)
	c.Bind(&testCategories)
	testUpdateCategories, err1 := database.UpdateTestCategories(testCategories)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}

	//custom data for body response
	output := map[string]interface{}{
		"CreatedAt":            testUpdateCategories.CreatedAt,
		"UpdatedAt":            testUpdateCategories.UpdatedAt,
		"DeletedAt":            testUpdateCategories.DeletedAt,
		"id":                   testUpdateCategories.ID,
		"test_categories_name": testUpdateCategories.Test_categories_Name,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                "success update test categories",
		"update test categories": output,
	})
}
