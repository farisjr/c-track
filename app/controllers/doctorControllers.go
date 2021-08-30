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
	//doctor.Role = "doctor"
	c.Bind(&input)
	//adddoctor, err := database.CreateDoctor(models.Doctor{})
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
	//doctor.Email = input.Email
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
		"ID": doctor.UserID,
		// "Name":  doctorAdd.Nama,
		// "Email": doctorAdd.Email,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new doctor",
		"customer": mapDoctor,
	})
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": "cannot insert data",
	// 	})
	// }
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "success create new doctor",
	// 	"data":    adddoctor,
	// })
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
		"doctor":  mapLogindoctor,
	})
}

//Authorization doctor
func DoctorAutorize(doctorId int, c echo.Context) error {
	authDoctor, err := database.GetOneDoctor(doctorId)
	LoggedInDoctorId, role := middlewares.ExtractTokenUserId(c)
	if LoggedInDoctorId != doctorId || string(authDoctor.Role) != role || err != nil || authDoctor.Role != "Doctor" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access")
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

// func GetDoctorsController(c echo.Context) error {
// 	doctors, err := database.GetDoctor()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get doctors data",
// 		"data":    doctors,
// 	})
// }

// func GetDoctorsIdController(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}
// 	doctors, err := database.GetDoctorById(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "cannot fetch data",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get doctor  by id",
// 		"data":    doctors,
// 	})
// }

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

// func UpdateDoctorsController(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}
// 	if err = DoctorAutorize(id, c); err != nil {
// 		return err
// 	}
// 	updateDoctors, _ := database.GetUpdateDoctor(id)
// 	c.Bind(&updateDoctors)
// 	testsId, err1 := database.UpdateDoctor(updateDoctors)
// 	if err1 != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success update doctor ",
// 		" data":   testsId,
// 	})
// }
