package routes

import (
	"app/constants"
	"app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	// e.POST("/register", controllers.UserRegister)
	//e.GET("/users/:id", controllers.GetUserById)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//------------------Non Authorized Test ----------------------//
	eJwt.POST("/tests", controllers.CreateTest)
	//eJwt.GET("/tests", controllers.GetTestsController)
	eJwt.GET("/tests/:id", controllers.GetOneTestController)
	eJwt.PUT("/tests/:id", controllers.UpdateTest)

	//------------------Non Authorized Checker ----------------------//
	// eJwt.GET("/checkers", controllers.GetCheckersController)
	// eJwt.GET("/checkers/:id", controllers.GetCheckerController)
	// eJwt.POST("/checkers", controllers.CreateCheckersController)

	//------------------Non Authorized Doctor ----------------------//
	// eJwt.GET("/doctors", controllers.GetDoctorsController)
	// eJwt.GET("/doctors/:id", controllers.GetDoctorsIdController)
	// eJwt.POST("/doctors", controllers.CreateDoctorsController)
	// eJwt.PUT("/doctors/:id", controllers.UpdateDoctorsController)

	///------------------Non Authorized Patient ----------------------//
	// eJwt.POST("/patients", controllers.CreatePatientsController)
	// eJwt.GET("/patients", controllers.GetPatientsController)
	// eJwt.GET("/patients/:id", controllers.GetPatientsIdController)
	// eJwt.PUT("/patients/:id", controllers.UpdatePatientsController)

	//------------------ Logout Controllers ----------------------//
	//eJwt.POST("/checkers/logout/:id", controllers.LogoutChecker)
	//eJwt.POST("/doctors/logout/:id", controllers.LogoutDoctor)
	// eJwt.POST("/patients/logout/:id", controllers.LogoutPatient)
	//eJwt.POST("/users/logout/:id", controllers.LogoutUserController)

}
