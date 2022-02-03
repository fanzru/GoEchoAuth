package controllers

import (
	"GoEchoAuth/config"
	"GoEchoAuth/models"
	"GoEchoAuth/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {

	data := &models.Users{}
	// bind data
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ResponseError(400, false, "Bind data Error!"))
	}

	// construct data
	user := &models.Users{
		Email:    data.Email,
		Password: data.Password,
	}

	// get connection
	db, err := config.ConnectionDatabase()
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ResponseError(400, false, "Connection Database Error!"))
	}

	err = db.Create(user).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ResponseError(400, false, "Input data failed!"))
	}

	return c.JSON(http.StatusOK, utils.ResponseSuccess(200, true, "Register Successfully ", user))
}
