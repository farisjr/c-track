package routes

import (
	"app/constants"
	"app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitAuth(e *echo.Echo) {
	e.POST("/login", controllers.LoginUserController)
	e.POST("/register", controllers.RegisterUserController)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	e.POST("/logout/:id", controllers.LogoutUserController)
}
