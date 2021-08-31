package routes

import (
	"app/constants"
	"app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	//------------------ Patient Login & Register Routes ----------------------//
	e.POST("/patient/login", controllers.PatientLogin)
	e.POST("/patient/register", controllers.PatientSignUp)

	//------------------ Doctor Login & Register Routes ----------------------//
	e.POST("/doctor/login", controllers.DoctorLogin)
	e.POST("/doctor/register", controllers.DoctorSignUp)

	//------------------ Checker Login & Register Routes ----------------------//
	e.POST("/checker/login", controllers.CheckerLogin)
	e.POST("/checker/register", controllers.CheckerSignUp)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	eJwt.POST("/doctor/test", controllers.DoctorCreateNewTest)      // doctor add new test routes
	eJwt.PUT("/doctor/test/:test_id", controllers.DoctorUpdateTest) // doctor update test result routes
	eJwt.GET("/doctor/tests/:doctor_id", controllers.DoctorGetTest) // doctor get all test

	eJwt.GET("/checker/test/:patient_id", controllers.CheckerGetTest) // checker get test by patient id

	eJwt.GET("/patient/test/:patient_id", controllers.PatientGetTest) //get patient test only himself

	//------------------ Logout Routes ----------------------//
	eJwt.POST("/checker/logout/:user_id", controllers.LogoutChecker)
	eJwt.POST("/doctor/logout/:user_id", controllers.LogoutDoctor)
	eJwt.POST("/patient/logout/:user_id", controllers.LogoutPatient)

}
