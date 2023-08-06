package crud

import (
	"github.com/tomp332/gospray/src"
	"github.com/tomp332/gospray/src/models"
	"github.com/tomp332/gospray/src/utils"
	"log"
)

type ICredentialsCrud struct{}

var CredentialsCrud = &ICredentialsCrud{}

// Get gets the credentials with the given id
func (c *ICredentialsCrud) Get(limit, offset uint) ([]models.IReadCredentials, error) {
	var fetchedCreds []models.CredentialsDTO
	err := src.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&fetchedCreds).Error
	if err != nil {
		return nil, err
	}
	return utils.TransformDTOCredentials(&fetchedCreds), nil

}

// Add adds the given credentials to the database
func (c *ICredentialsCrud) Add(creds []models.ICredentialsCreate) ([]models.IReadCredentials, error) {
	credsModelSlice := make([]models.CredentialsDTO, len(creds))
	for i, credsBase := range creds {
		err := utils.CopyStructFields(credsBase, &credsModelSlice[i])
		if err != nil {
			return nil, err
		}
	}
	result := src.MainDB.Create(credsModelSlice)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Added new credentials to DB, count: %d", len(credsModelSlice))
	return utils.TransformDTOCredentials(&credsModelSlice), nil
}

// Update updates the credentials with the given id
func (c *ICredentialsCrud) Update(creds []*models.IUpdateCredentials) ([]models.IReadCredentials, error) {
	updatedCredentials := make([]models.CredentialsDTO, len(creds))
	for i, updateSchema := range creds {
		updatedCredentials[i] = models.CredentialsDTO{
			CustomORMModel: models.CustomORMModel{
				ID: updateSchema.ID,
			},
			CredentialsBase: updateSchema.CredentialsBase,
		}
	}
	result := src.MainDB.Save(updatedCredentials)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Updated credentials information with id %d", updatedCredentials[0].ID)
	return utils.TransformDTOCredentials(&updatedCredentials), nil
}

// Delete deletes the credentials with the given id
func (c *ICredentialsCrud) Delete(id uint) error {
	result := src.MainDB.Delete(&models.CredentialsDTO{}, id)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Deleted credentials with id %d", id)
	return nil
}
