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
// @Param credentials body []models.ICredCreate true "Credentials"
// @Success 200 {array} models.CredsModel
// @Failure 400 {object} models.ServerError
// @Failure 500 {object} models.ServerError
// @Router /creds [post]
func AddCreds(c echo.Context) error {
	var creds []*models.CredsModel
	err := c.Bind(&creds)
	if err != nil || creds == nil {
		log.Printf("Error binding creds struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("Validation error for CredsModel schema", err))
	}
	addedCreds, err := crud.Add(creds)
	if err != nil {
		return c.JSONBlob(http.StatusInternalServerError,
			utils.BadRequestError("Error adding new credentials to database", err))
	}
	return c.JSON(http.StatusOK, addedCreds)
}

// GetCreds godoc
// @Summary Get credentials
// @Description Get credentials from the database
// @Tags Creds
// @Param limit query int false "Limit the number of results"
// @Param page query int false "Page number"
// @Success 200 {array} models.CredsModel
// @Failure 400 {object} models.ServerError
// @Failure 500 {object} models.ServerError
// @Router /creds [get]
func GetCreds(c echo.Context) error {
	paginationParams := new(src.Paginate)
	err := c.Bind(paginationParams)
	if err != nil || paginationParams == nil {
		log.Printf("Error binding paginate struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("One or more of the parameters specified for pagination was incorrect", err))
	}
	creds, err := crud.Get(paginationParams.Limit, paginationParams.Page, []*models.CredsModel{})
	if err != nil {
		log.Printf("Error getting creds: %s", err.Error())
		return c.JSONBlob(http.StatusInternalServerError,
			utils.BadRequestError("Error fetching credentials from the database", err))
	}
	return c.JSON(http.StatusOK, creds)
}
