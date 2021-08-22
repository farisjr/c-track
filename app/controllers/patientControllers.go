package controllers

import (
	"app/lib/database"
	"app/middlewares"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Register patient controller for patient registration
func RegisterPatientController(c echo.Context) error {
	patient := models.User{}
	patient.Role = "patient"
	c.Bind(&patient)
	addpatient, err := database.CreatePatient(models.Patient{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new patient",
		"data":    addpatient,
	})
}

//Login for patient with matching username and password
func LoginPatient(c echo.Context) error {
	patient := models.User{}
	c.Bind(&patient)
	loginpatient, err := database.PatientLoginDB(patient.Username, patient.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mapLoginpatient := map[string]interface{}{
		"UserID": loginpatient.ID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"doctor":  mapLoginpatient,
	})
}

//Authorization patient
func AuthorizationPatient(patientId int, c echo.Context) error {
	authpatient, err := database.GetPatientById(patientId)
	loggedInPatientId, role := middlewares.ExtractTokenUserId(c)
	if loggedInPatientId != patientId || string(authpatient.User.Role) != role || err != nil || authpatient.User.Role != "patient" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access")
	}
	return nil
}

//Logout patient
func LogoutPatient(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("patientId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, err := database.GetPatientById(id)
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

func GetPatientsController(c echo.Context) error {
	patients, err := database.GetPatient()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get patients data",
		"data":    patients,
	})
}

func GetPatientsIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	patient, err := database.GetPatientById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get patient by id",
		"data":    patient,
	})
}

func CreatePatientsController(c echo.Context) error {
	//binding data
	patient := models.Patient{}
	c.Bind(&patient)
	//Checking if id already exist
	duplicate := database.CheckDuplicatePatient(patient)
	if duplicate != nil {
		return echo.NewHTTPError(http.StatusBadRequest, duplicate.Error())
	}
	patients, err := database.CreatePatient(patient)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success Create Patient",
		"patient":  patients,
	})
}

func UpdatePatientsController(c echo.Context) error {
	var patient models.Patient
	// Validation of id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	// Getting Patient Data by id
	GetPatient, err := database.GetPatientById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	patient = GetPatient
	// Updating Patient Data
	c.Bind(&patient)
	update_patient, err := database.UpdatePatient(patient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success update patient profile",
		"patient": update_patient,
	})
}

func DeletePatientsController(c echo.Context) error {
	//Validation of id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	// Deleting Patient Data
	delete_patient, err := database.DeletePatient(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  delete_patient,
	})
}
