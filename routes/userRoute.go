package routes

import (
	"GoEchoAuth/controllers"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo) *echo.Echo {
	e.POST("/api/register", controllers.RegisterUser)
	e.POST("/api/login", controllers.LoginUser)
	return e
}
