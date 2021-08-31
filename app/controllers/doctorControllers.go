package controllers

import (
	"app/lib/database"
	"app/middlewares"
	"app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Register doctor controller for doctor registration
func DoctorSignUp(c echo.Context) error {
	input := models.User{}
	c.Bind(&input)
	if input.UserID == 0 || input.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please fill userid and password correctly",
		})
	}
	if same, _ := database.CheckSameUserId(input.UserID); same == true {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "userid already used",
		})
	}
	addDoctor := models.User{}
	addDoctor.UserID = input.UserID
	addDoctor.Password = ourEncrypt(input.Password)
	addDoctor.Role = "Doctor"
	c.Bind(&addDoctor)
	doctor, err := database.CreateDoctor(addDoctor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapDoctor := map[string]interface{}{
		"User ID": doctor.UserID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new doctor",
		"data":    mapDoctor,
	})
}

//Login for doctor with matching userid and password
func DoctorLogin(c echo.Context) error {
	input := models.User{}
	c.Bind(&input)
	logindoctor, err := database.DoctorLoginDB(input.UserID, input.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mapLogindoctor := map[string]interface{}{
		"UserID": logindoctor.UserID,
		"Token":  logindoctor.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes login",
		"data":    mapLogindoctor,
	})
}

//Authorization doctor
func DoctorAutorize(doctorId int, c echo.Context) error {
	authDoctor, err := database.GetOneDoctor(doctorId)
	LoggedInDoctorId, role := middlewares.ExtractTokenUserId(c)
	if LoggedInDoctorId != doctorId || string(authDoctor.Role) != role || err != nil || authDoctor.Role != "Doctor" {
		return echo.NewHTTPError(http.StatusUnauthorized, "This user does not have access")
	}
	return nil
}

//Logout doctor
func LogoutDoctor(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, _ := database.GetOneUser(userId)
	logout.Token = ""
	doctor, err := database.UpdateDoctor(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Thank you",
		"data":    doctor,
	})
}
