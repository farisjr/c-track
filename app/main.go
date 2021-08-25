package main

import (
	"app/config"
	"app/middlewares"
	"app/routes"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDb()
	config.InitPort()
	middlewares.LogMiddlewares((e))
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
