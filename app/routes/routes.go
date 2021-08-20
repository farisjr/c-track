package routes

import (
	"app/controllers"

	"github.com/labstack/echo"
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

}
