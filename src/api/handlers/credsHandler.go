package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src/crud"
	"github.com/tomp332/bruteForcer/src/models"
	"net/http"
	"strconv"
)

// AddCreds godoc
// @Summary Add credentials
// @Description Add credentials to the database
// @Tags Creds
// @Accept json
// @Produce json
// @Success 200 {JSON} JSON with the added credentials
// @Failure 400 {JSON} JSON Bad Request
// @Router /creds [post]
func AddCreds(c echo.Context) error {
	var creds []*models.Cred
	err := c.Bind(&creds)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(http.StatusOK, creds)
}

// GetCreds godoc
// @Summary Get credentials
// @Description Get credentials from the database
// @Tags Creds
// @Accept json
// @Produce json
// @Success 200 {string} JSON with all credentials
// @Failure 400 {string} string "Bad Request"
// @Router /creds [get]
func GetCreds(c echo.Context) error {
	limit := c.QueryParam("limit")
	creds, err := crud.GetCreds(strconv.Atoi(limit))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, creds)
}
