package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src/crud"
	"github.com/tomp332/bruteForcer/src/models"
	"log"
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
	creds, err = crud.AddCreds(creds)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(http.StatusOK, creds)
}

// GetCreds godoc
// @Summary Get credentials
// @Description Get credentials from the database
// @Tags Creds
// @Produce json
// @Param limit query int false "Limit the number of results"
// @Param page query int false "Page number"
// @Success 200 {JSON} JSON with the credentials
// @Failure 400 {JSON} JSON Bad Request
// @Router /creds [get]
func GetCreds(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	limitInt, err := strconv.Atoi(limit)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Printf("Error parsing limit or page: %e", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	creds, err := crud.GetCreds(limitInt, pageInt)
	if err != nil {
		log.Printf("Error getting creds: %e", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, creds)
}
