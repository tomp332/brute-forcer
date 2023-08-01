package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/crud"
	"github.com/tomp332/bruteForcer/src/models"
	"github.com/tomp332/bruteForcer/src/utils"
	"log"
	"net/http"
)

// AddCreds godoc
// @Summary Add credentials
// @Description Add credentials to the database
// @Tags Creds
// @Accept json
// @Success 200 {array} models.Cred
// @Failure 400 {object} models.ServerError
// @Failure 500 {object} models.ServerError
// @Router /creds [post]
func AddCreds(c echo.Context) error {
	var creds []*models.Cred
	err := c.Bind(&creds)
	if err != nil || creds == nil {
		log.Printf("Error binding creds struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("Validation error for Cred schema", err))
	}
	creds, err = crud.AddCreds(creds)
	if err != nil {
		return c.JSONBlob(http.StatusInternalServerError,
			utils.BadRequestError("Error adding new credentials to database", err))
	}
	return c.JSON(http.StatusOK, creds)
}

// GetCreds godoc
// @Summary Get credentials
// @Description Get credentials from the database
// @Tags Creds
// @Param limit query int false "Limit the number of results"
// @Param page query int false "Page number"
// @Success 200 {array} models.Cred
// @Failure 400 {object} models.ServerError
// @Failure 500 {object} models.ServerError
// @Router /creds [get]
func GetCreds(c echo.Context) error {
	var paginateStruct *src.Paginate
	err := c.Bind(&paginateStruct)
	if err != nil || paginateStruct == nil {
		log.Printf("Error binding paginate struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("One or more of the parameters specified for pagination was incorrect", err))
	}
	creds, err := crud.GetCreds(paginateStruct.Limit, paginateStruct.Page)
	if err != nil {
		log.Printf("Error getting creds: %s", err.Error())
		return c.JSONBlob(http.StatusInternalServerError,
			utils.BadRequestError("Error fetching credentials from the database", err))
	}
	return c.JSON(http.StatusOK, creds)
}
