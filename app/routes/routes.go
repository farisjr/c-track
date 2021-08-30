package routes

import (
	"app/constants"
	"app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	eJwt.POST("/doctor/test", controllers.DoctorCreateNewTest)      // doctor add new test routes
	eJwt.PUT("/doctor/test/:test_id", controllers.DoctorUpdateTest) // doctor update test result routes
	eJwt.GET("/doctor/tests", controllers.DoctorGetAllTest)         // doctor get all test

	eJwt.GET("/checker/test/:patient_id", controllers.CheckerGetTest) // checker get test by patient id

	//------------------ Logout Controllers ----------------------//
	//eJwt.POST("/checkers/logout/:user_id", controllers.LogoutChecker)
	eJwt.POST("/doctors/logout/:user_id", controllers.LogoutDoctor)
	//eJwt.POST("/patients/logout/:user_id", controllers.LogoutPatient)

}
