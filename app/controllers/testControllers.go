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

func AuthorizedChecker(c echo.Context) bool {
	_, role := middlewares.ExtractTokenUserId(c)
	if role != "Checker" {
		return false
	}
	return true
}

func AuthorizedPatient(c echo.Context) bool {
	_, role := middlewares.ExtractTokenUserId(c)
	if role != "Patient" {
		return false
	}
	return true
}

//Doctor feature for create new test
func DoctorCreateNewTest(c echo.Context) error {
	auth := AuthorizedDoctor(c)
	if !auth {
		return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
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

//Doctor feature for get all test data
func DoctorGetAllTest(c echo.Context) error {
	auth := AuthorizedDoctor(c)
	if !auth {
		return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	}
	tests, _ := database.GetAllTests()
	if len(tests) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "No test data")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all test data",
		"data":    tests,
	})
}

//Doctor feature for updating test result
func DoctorUpdateTest(c echo.Context) error {
	auth := AuthorizedDoctor(c)
	if !auth {
		return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
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

//Checker feature for get test by patient id
func CheckerGetTest(c echo.Context) error {
	auth := AuthorizedChecker(c)
	if !auth {
		return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access to this route")
	}
	patientId, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	tests, err := database.GetTestbyPatient(patientId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get test by patient id",
		"data":    tests,
	})
}

//Patient feature for get all test
func PatientGetTest(c echo.Context) error {
	patientId, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	if err = PatientAuthorize(patientId, c); err != nil {
		return err
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Record not found",
		})
	}
	tests, err := database.GetTestbyPatient(patientId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all patient test  result",
		"data":    tests,
	})
}

//Doctor get test by doctor id
func DoctorGetTest(c echo.Context) error {
	doctorId, err := strconv.Atoi(c.Param("doctor_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	if err = DoctorAutorize(doctorId, c); err != nil {
		return err
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Record not found",
		})
	}
	tests, err := database.GetTestbyDoctor(doctorId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all patient test  result",
		"data":    tests,
	})
}
