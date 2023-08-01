package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/crud"
	"github.com/tomp332/bruteForcer/src/models"
	"log"
	"net/http"
)

// AddSlaves godoc (POST /slaves)
// @Summary Add slaves
// @Description Add slaves to the database
// @Tags Slaves
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
// @Summary Get slaves
// @Description Get slaves from the database
// @Tags Slaves
// @Produce json
// @Param limit query int false "Limit the number of results"
// @Param page query int false "Page number"
// @Success 200 {array} models.Slave
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /slaves [get]
func GetSlaves(c echo.Context) error {
	var paginateStruct *src.Paginate
	err := c.Bind(&paginateStruct)
	if err != nil || paginateStruct == nil {
		log.Printf("Error binding paginate struct: %e", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	slaves, err := crud.GetSlaves(paginateStruct.Limit, paginateStruct.Page)
	if err != nil {
		log.Printf("Error getting slaves: %e", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, slaves)
}
