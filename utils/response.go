package utils

import (
	"github.com/labstack/echo/v4"
)

func ResponseError(code int, status bool, message string) echo.Map {
	return echo.Map{
		"code":    code,
		"status":  status,
		"message": message,
	}
}

func ResponseSuccess(code int, status bool, message string, data interface{}) echo.Map {
	return echo.Map{
		"code":    code,
		"status":  status,
		"message": message,
		"data":    data,
	}
}
