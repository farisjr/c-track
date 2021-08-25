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
func RegisterDoctorController(c echo.Context) error {
	doctor := models.User{}
	doctor.Role = "doctor"
	c.Bind(&doctor)
	adddoctor, err := database.CreateDoctor(models.Doctor{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new doctor",
		"data":    adddoctor,
	})
}

//Login for doctor with matching username and password
func LoginDoctor(c echo.Context) error {
	doctor := models.User{}
	c.Bind(&doctor)
	logindoctor, err := database.LoginDoctorDB(doctor.Username, doctor.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mapLogindoctor := map[string]interface{}{
		"UserID": logindoctor.ID,
		"Name":   logindoctor.Name,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes login",
		"doctor":  mapLogindoctor,
	})
}

//Authorization doctor
func AuthorizationDoctor(doctorId int, c echo.Context) error {
	authdoctor, err := database.GetDoctorById(doctorId)
	LoggedInDoctorId, role := middlewares.ExtractTokenUserId(c)
	if LoggedInDoctorId != doctorId || string(authdoctor.User.Role) != role || err != nil || authdoctor.User.Role != "doctor" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access")
	}
	return nil
}

//Logout doctor
func LogoutDoctor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("doctorId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, err := database.GetDoctorById(id)
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

// func DeleteDoctorsController(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}
// 	doctors, _ := database.GetDoctorById(id)
// 	c.Bind(&doctors)
// 	doctorsDeleted, err := database.DeleteDoctorById(doctors)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "cannot delete data",
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success delete selected doctor",
// 		"data":    doctorsDeleted,
// 	})
// }

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
