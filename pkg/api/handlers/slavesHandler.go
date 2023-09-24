package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/gobrute/pkg/crud"
	"github.com/tomp332/gobrute/pkg/types"
	"github.com/tomp332/gobrute/pkg/utils"

	"log"
	"net/http"
)

// AddSlaves godoc (POST /slaves)
// @Summary Add slaves
// @Description Add slaves to the database
// @Tags Slaves
// @Accept json
// @Produce json
// @Success 200 {array} managerTypes.SlaveDTO
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /slaves [post]
func AddSlaves(c echo.Context) error {
	var slaves []*types.SlaveDTO
	err := c.Bind(&slaves)
	if err != nil || slaves == nil {
		log.Printf("Error binding slaves struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("Validation error for SlavesdModel schema", err))
	}
	newSlaves, err := crud.SlavesCrud.Add(slaves)
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
// @Param page query int false "Offset number"
// @Success 200 {array} managerTypes.SlaveDTO
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /slaves [get]
func GetSlaves(c echo.Context) error {
	paginationParams := new(crud.IPaginateParams)
	err := c.Bind(paginationParams)
	if err != nil || paginationParams == nil {
		log.Printf("Error binding pagination struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("One or more of the parameters specified for pagination was incorrect", err))
	}
	creds, err := crud.SlavesCrud.Get(paginationParams.Limit, paginationParams.Offset)
	if err != nil {
		log.Printf("Error getting creds: %s", err.Error())
		return c.JSONBlob(http.StatusInternalServerError,
			utils.BadRequestError("Error fetching credentials from the database", err))
	}
	return c.JSON(http.StatusOK, creds)
}
