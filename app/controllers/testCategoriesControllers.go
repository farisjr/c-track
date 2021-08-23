package controllers

import (
	"app/config"
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

func GetTestCategoriesController(c echo.Context) error {
	var testCategories []models.TestCategories
	err := config.DB.Debug().Model(&models.TestCategories{}).Find(&testCategories).Error
	//testCategories, err := database.GetTestCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.TestCategoriesResponse{
			false, "Failed get database test category", nil,
		})
	}
	return c.JSON(http.StatusOK, models.TestCategoriesResponse{
		true, "Success", testCategories,
	})
}

func GetTestCategoriesIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	testCategories, err := database.GetTestCategory(id)
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

func DeleteTestCategoriesByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	testCategories, err := database.GetTestCategory(id)
	c.Bind(&testCategories)
	testCategoriesDeleted, err := database.DeleteTestCategory(testCategories)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot delete data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete success ",
		"data":    testCategoriesDeleted,
	})
}

func UpdateTestCategoriesController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	testCategories, err := database.GetTestCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get data",
		})
	}
	c.Bind(&testCategories)
	testUpdateCategories, err := database.UpdateTestCategory(testCategories)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                "success update test categories",
		"update test categories": testUpdateCategories,
	})
}
