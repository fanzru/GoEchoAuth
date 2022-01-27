package main

import (
	"GoEchoAuth/config"
	"GoEchoAuth/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Connected struct {
	Status  bool   `json:"status" xml:"status"`
	Code    string `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func main() {
	_, err := config.ConnectionDatabase()
	if err != nil {
		panic(err.Error())
	}
	e := routes.Init()

	e.GET("/", func(c echo.Context) error {
		C := &Connected{
			Status:  true,
			Code:    "200",
			Message: "Api Connected successfully",
		}
		return c.JSON(http.StatusOK, C)
	})

	e.Logger.Fatal(e.Start(":1002"))

}
