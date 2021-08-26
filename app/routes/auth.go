package routes

import (
	"app/controllers"

	"github.com/labstack/echo"
)

func InitAuth(e *echo.Echo) {
	e.POST("/login", controllers.LoginUserController)
	e.POST("/register", controllers.RegisterUserController)
}
