package controllers

import (
	"app/lib/database"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateDoctorsController(c echo.Context) error {
	doctors := models.Doctor{}
	c.Bind(&doctors)

	doctorsAdd, err := database.CreateDoctor(doctors)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new doctor",
		"data":    doctorsAdd,
	})
}

func GetDoctorsController(c echo.Context) error {
	doctors, err := database.GetDoctor()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get doctors data",
		"data":    doctors,
	})
}

func GetDoctorsIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	doctors, err := database.GetDoctorById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get doctor  by id",
		"data":    doctors,
	})
}

func DeleteDoctorsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	doctors, _ := database.GetDoctorById(id)
	c.Bind(&doctors)
	doctorsDeleted, err := database.DeleteDoctorById(doctors)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot delete data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete selected doctor",
		"data":    doctorsDeleted,
	})
}

func UpdateDoctorsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	updateDoctors := database.GetUpdateDoctor(id)
	c.Bind(&updateDoctors)
	testsId, err1 := database.UpdateDoctor(updateDoctors)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update doctor ",
		" data":   testsId,
	})
}
