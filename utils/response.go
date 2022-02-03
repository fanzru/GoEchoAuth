package utils

import "github.com/labstack/echo/v4"

func ResponseError(code int, status bool, message string) echo.Map {
	return echo.Map{
		"Code":    code,
		"Status":  status,
		"Message": message,
	}
}

func ResponseSuccess(code int, status bool, message string, data interface{}) echo.Map {
	return echo.Map{
		"Code":    code,
		"Status":  status,
		"Message": message,
		"Data":    data,
	}
}
