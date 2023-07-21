package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src/crud"
	"github.com/tomp332/bruteForcer/src/models"
	"net/http"
)

func AddSlaves(c echo.Context) error {
	var slaves []*models.Slave
	err := c.Bind(&slaves)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newSlaves, err := crud.AddSlaves(slaves)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, newSlaves)
}

func GetSlaves(c echo.Context) error {
	slaves, err := crud.GetSlaves()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, slaves)
}
