package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src/crud"
	"github.com/tomp332/bruteForcer/src/models"
	"net/http"
)

// AddSlaves godoc (POST /slaves)
// @Summary Add slaves
// @Description Add slaves to the database
// @Tags slaves
// @Accept json
// @Produce json
// @Success 200 {array} models.Slave
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /slaves [post]
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

// GetSlaves ... Get all slaves
// @Summary Get all slaves
// @Description get all users
// @Tags Users
// @Success 200 {JSON} model.Slave
// @Failure 404 {object} object
// @Router / [get]
func GetSlaves(c echo.Context) error {
	slaves, err := crud.GetSlaves()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, slaves)
}
