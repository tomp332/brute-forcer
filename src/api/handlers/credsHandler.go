package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/crud"
	"github.com/tomp332/bruteForcer/src/models"
	"log"
	"net/http"
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
	var paginateStruct *src.Paginate
	err := c.Bind(&paginateStruct)
	if err != nil {
		log.Printf("Error binding paginate struct: %e", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	creds, err := crud.GetCreds(paginateStruct.Limit, paginateStruct.Page)
	if err != nil {
		log.Printf("Error getting creds: %e", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, creds)
}
