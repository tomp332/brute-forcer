package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"github.com/tomp332/gobrute/pkg/manager"
	crud2 "github.com/tomp332/gobrute/pkg/manager/crud"
	"github.com/tomp332/gobrute/pkg/manager/utils"

	"log"
	"net/http"
)

// AddCreds godoc
// @Summary Add credentials
// @Description Add credentials to the database
// @Tags Creds
// @Accept json
// @Param credentials body []internalTypes.ICredentialsCreate true "ICredentialsCreate"
// @Success 200 {array} internalTypes.ICredentialsCreate
// @Failure 400 {object} internalTypes.ServerError
// @Failure 500 {object} internalTypes.ServerError
// @Router /creds [post]
func AddCreds(c echo.Context) error {
	var creds []internalTypes.ICredentialsCreate
	err := c.Bind(&creds)
	if err != nil || creds == nil {
		log.Printf("Error binding creds struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("Validation error for CredentialsDTO schema", err))
	}
	if err != nil {
		log.Printf("Error validating creds struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("Validation error for CredentialsDTO schema", err))
	}
	addedCreds, err := crud2.CredentialsCrud.Add(creds)
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
// @Param page query int false "Offset number"
// @Success 200 {array} internalTypes.IReadCredentials
// @Failure 400 {object} internalTypes.ServerError
// @Failure 500 {object} internalTypes.ServerError
// @Router /creds [get]
func GetCreds(c echo.Context) error {
	paginationParams := new(crud2.IPaginateParams)
	err := c.Bind(paginationParams)
	if paginationParams.Limit == 0 {
		// Specify default values for pagination
		paginationParams.Limit = manager.GlobalSettings.APISettings.PaginationDefaultLimit
	}
	if err != nil || paginationParams == nil {
		log.Printf("Error binding pagination struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("One or more of the parameters specified for pagination was incorrect", err))
	}
	creds, err := crud2.CredentialsCrud.Get(paginationParams.Limit, paginationParams.Offset)
	if err != nil {
		log.Printf("Error getting credentials: %s", err.Error())
		return c.JSONBlob(http.StatusInternalServerError,
			utils.BadRequestError("Error fetching credentials from the database", err))
	}
	return c.JSON(http.StatusOK, creds)
}

// UpdateCreds godoc
// @Summary Update credentials
// @Description Update credentials in the database
// @Tags Creds
// @Accept json
// @Param credentials body []internalTypes.IUpdateCredentials true "ICredentialsCreate"
// @Success 200 {array} internalTypes.IReadCredentials
// @Failure 400 {object} internalTypes.ServerError
// @Failure 500 {object} internalTypes.ServerError
// @Router /creds [put]
func UpdateCreds(c echo.Context) error {
	var creds []*internalTypes.IUpdateCredentials
	err := c.Bind(&creds)
	if err != nil || creds == nil {
		log.Printf("Error binding credentials struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("Validation error for CredentialsDTO schema", err))
	}
	updatedCreds, err := crud2.CredentialsCrud.Update(creds)
	if err != nil {
		return c.JSONBlob(http.StatusInternalServerError,
			utils.BadRequestError("Error updating credentials in database", err))
	}
	return c.JSON(http.StatusOK, updatedCreds)
}

// DeleteCreds godoc
// @Summary Delete credentials by ID
// @Description IDeleteParams credentials from the database
// @Tags Creds
// @Param id query string true "ID of the credentials to delete"
// @Success 200
// @Failure 400 {object} internalTypes.ServerError
// @Failure 500 {object} internalTypes.ServerError
// @Router /creds [delete]
func DeleteCreds(c echo.Context) error {
	deleteParams := new(crud2.IDeleteParams)
	err := c.Bind(deleteParams)
	if err != nil || deleteParams == nil {
		log.Printf("Error binding delete struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("One or more of the parameters specified for delete was incorrect", err))
	}
	err = crud2.CredentialsCrud.Delete(deleteParams.ID)
	if err != nil {
		log.Printf("Error deleting credentials: %s", err.Error())
		return c.JSONBlob(http.StatusInternalServerError,
			utils.BadRequestError("Error deleting credentials from the database", err))
	}
	return c.JSON(http.StatusOK, nil)
}
