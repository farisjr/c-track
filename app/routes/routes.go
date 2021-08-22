package routes

import (
	"app/constants"
	"app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	//------------------Non Authorized Test Categories----------------------//
	e.GET("/testCategories", controllers.GetTestCategoriesController)
	e.GET("/testCategories/:id", controllers.GetTestCategoriesIdController)
	e.POST("/testCategories", controllers.CreateTestCategoriesController)
	e.PUT("/testCategories/:id", controllers.UpdateTestCategoriesController)
	e.DELETE("/testCategories/:id", controllers.DeleteTestCategoriesByIdController)

	//------------------Non Authorized Test ----------------------//
	e.GET("/tests", controllers.GetTestsController)
	e.GET("/tests/:id", controllers.GetTestsIdController)
	e.POST("/tests", controllers.CreateTestsController)
	e.PUT("/tests/:id", controllers.UpdateTestsController)
	e.DELETE("/tests/:id", controllers.DeleteTestsController)

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

	//AUTHORIZATION JWT
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

}
