package controllers

import (
	"app/lib/database"
	"app/middlewares"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Register patient feature for patient registration
func PatientSignUp(c echo.Context) error {
	input := models.User{}
	c.Bind(&input)
	if input.UserID == 0 || input.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please fill userid and password correctly",
		})
	}
	if same, _ := database.CheckSameId(input.UserID); same == true {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "userid already used",
		})
	}
	addPatient := models.User{}
	addPatient.UserID = input.UserID
	addPatient.Password = ourEncrypt(input.Password)
	addPatient.Role = "Patient"
	c.Bind(&addPatient)
	patient, err := database.CreatePatient(addPatient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapPatient := map[string]interface{}{
		"User ID": patient.UserID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new patient",
		"data":    mapPatient,
	})
}

//Login for patient with matching userid and password
func PatientLogin(c echo.Context) error {
	patient := models.User{}
	c.Bind(&patient)
	loginpatient, err := database.PatientLoginDB(patient.UserID, patient.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mapLoginpatient := map[string]interface{}{
		"User ID": loginpatient.UserID,
		"Token":   loginpatient.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    mapLoginpatient,
	})
}

//Authorization patient
func PatientAuthorize(userId int, c echo.Context) error {
	patientAuth, err := database.GetOnePatient(userId)
	loggedInPatientId, role := middlewares.ExtractTokenUserId(c)
	if loggedInPatientId != userId || string(patientAuth.Role) != role || err != nil || patientAuth.Role != "Patient" {
		return echo.NewHTTPError(http.StatusUnauthorized, "This account does not have access")
	}
	return nil
}

//Logout patient
func LogoutPatient(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, _ := database.GetOnePatient(userId)
	logout.Token = ""
	patient, err := database.UpdatePatient(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Thank you",
		"data":    patient,
	})
}
