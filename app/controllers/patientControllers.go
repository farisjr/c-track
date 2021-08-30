package controllers

import (
	"app/lib/database"
	"app/middlewares"
	"app/models"
	"net/http"

	"github.com/labstack/echo"
)

//Register patient controller for patient registration
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
		"patient": mapLoginpatient,
	})
}

//Authorization patient
func AuthorizationPatient(patientId int, c echo.Context) error {
	authpatient, err := database.GetOnePatient(patientId)
	loggedInPatientId, role := middlewares.ExtractTokenUserId(c)
	if loggedInPatientId != patientId || string(authpatient.Role) != role || err != nil || authpatient.Role != "Patient" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access")
	}
	return nil
}

//Logout patient
// func LogoutPatient(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("patientId"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}
// 	logout, _ := database.GetOnePatient(id)
// 	logout.Token = ""
// 	patient, err := database.UpdatePatient(logout)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "cannot logout",
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Thank you",
// 		"data":    patient,
// 	})
// }

// func GetPatientsController(c echo.Context) error {
// 	patients, err := database.GetPatient()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get patients data",
// 		"data":    patients,
// 	})
// }

// func GetPatientsIdController(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}
// 	patient, err := database.GetPatientById(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "cannot fetch data",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get patient by id",
// 		"data":    patient,
// 	})
// }

// func CreatePatientsController(c echo.Context) error {
// 	//binding data
// 	patient := models.Patient{}
// 	c.Bind(&patient)
// 	//Checking if id already exist
// 	patients, err := database.CreatePatient(patient)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest)
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"messages": "Success Create Patient",
// 		"patient":  patients,
// 	})
// }

// func UpdatePatientsController(c echo.Context) error {
// 	var patient models.Patient
// 	// Validation of id
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}
// 	// Getting Patient Data by id
// 	GetPatient, err := database.GetPatientById(id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest)
// 	}
// 	patient = GetPatient
// 	// Updating Patient Data
// 	c.Bind(&patient)
// 	update_patient, err := database.UpdatePatient(patient)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "can not fetch data",
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status":  "success update patient profile",
// 		"patient": update_patient,
// 	})
// }

// func DeletePatientsController(c echo.Context) error {
// 	//Validation of id
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}
// 	// Deleting Patient Data
// 	delete_patient, err := database.DeletePatient(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "can not fetch data",
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": "success",
// 		"users":  delete_patient,
// 	})
// }
