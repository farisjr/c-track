package routes

import (
	"app/constants"
	"app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	//------------------Non Authorized Test Categories----------------------//
	// e.GET("/testCategories", controllers.GetTestCategoriesController)
	// e.GET("/testCategories/:id", controllers.GetTestCategoriesIdController)
	// e.POST("/testCategories", controllers.CreateTestCategoriesController)
	// e.PUT("/testCategories/:id", controllers.UpdateTestCategoriesController)
	// e.DELETE("/testCategories/:id", controllers.DeleteTestCategoriesByIdController)

	//------------------Non Authorized Test ----------------------//
	e.GET("/tests", controllers.GetTestsController)
	e.GET("/tests/:id", controllers.GetTestsIdController)
	e.POST("/tests", controllers.CreateTestsController)
	e.PUT("/tests/:id", controllers.UpdateTestsController)
	e.DELETE("/tests/:id", controllers.DeleteTestsController)

	//------------------Non Authorized User ----------------------//
	e.POST("/users", controllers.RegisterUserController)

	//------------------Non Authorized Checker ----------------------//
	e.GET("/checkers", controllers.GetCheckerController)
	e.GET("/checkers/:id", controllers.GetCheckerController)
	e.POST("/checkers", controllers.CreateCheckersController)

	//------------------Non Authorized Doctor ----------------------//
	e.GET("/doctors", controllers.GetDoctorsController)
	e.GET("/doctors/:id", controllers.GetDoctorsIdController)
	e.POST("/doctors", controllers.CreateDoctorsController)
	e.PUT("/doctors/:id", controllers.UpdateDoctorsController)
	e.DELETE("/doctors/:id", controllers.DeleteDoctorsController)

	///------------------Non Authorized Patient ----------------------//
	e.GET("/patients", controllers.GetPatientsController)
	e.GET("/patients/:id", controllers.GetPatientsIdController)
	e.POST("/patients", controllers.CreatePatientsController)
	e.PUT("/patients/:id", controllers.UpdatePatientsController)
	e.DELETE("/patients/:id", controllers.DeletePatientsController)

	///------------------Non Authorized users ----------------------//
	// e.GET("/users", controllers.GetPatientsController)
	// e.GET("/users/:id", controllers.GetPatientsIdController)
	// e.POST("/users", controllers.CreatePatientsController)
	// e.PUT("/users/:id", controllers.UpdatePatientsController)
	// e.DELETE("/users/:id", controllers.DeletePatientsController)

	//------------------ Login Controllers ----------------------//
	e.POST("/checkers/login", controllers.LoginChecker)
	e.POST("/doctors/login", controllers.LoginDoctor)
	e.POST("/patients/login", controllers.LoginPatient)
	e.POST("/users/login", controllers.LoginUserController)
	e.POST("/register", controllers.RegisterUserController)

	//------------------ Logout Controllers ----------------------//
	e.POST("/checkers/logout/:id", controllers.LogoutChecker)
	e.POST("/doctors/logout/:id", controllers.LogoutDoctor)
	e.POST("/patients/logout/:id", controllers.LogoutPatient)
	e.POST("/users/logout/:id", controllers.LogoutUserController)

	//AUTHORIZATION JWT
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//------------------ Adding, Updating and Deleting Covid Test for Doctors ----------------------//

	//------------------ Login Controllers ----------------------//
	eJwt.POST("/checkers/login", controllers.LoginChecker)
	eJwt.POST("/doctors/login", controllers.LoginDoctor)
	eJwt.POST("/patients/login", controllers.LoginPatient)
	eJwt.POST("/users/login", controllers.LoginUserController)

	//------------------ Creating Covid Test for Doctors ----------------------//

	eJwt.POST("/tests", controllers.CreateTestsController)

	//------------------ Getting Covid Test for Checkers ----------------------//
	eJwt.GET("/tests/:id", controllers.GetTestsIdController)

}
