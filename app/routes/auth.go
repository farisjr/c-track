package routes

import (
	"app/constants"
	"app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitAuth(e *echo.Echo) {
	//e.POST("/login", controllers.LoginPatientController)
	//e.POST("/register", controllers.RegisterUserController)

	e.POST("/patient/login", controllers.PatientLogin)
	e.POST("/patient/register", controllers.PatientSignUp)

	e.POST("/doctor/login", controllers.DoctorLogin)
	e.POST("/doctor/register", controllers.DoctorSignUp)

	e.POST("/checker/login", controllers.CheckerLogin)
	e.POST("/checker/register", controllers.CheckerSignUp)

	e.GET("/tests/:id", controllers.GetOneTestController)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	//e.POST("/logout/:id", controllers.LogoutUserController)
}
