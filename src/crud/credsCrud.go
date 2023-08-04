package crud

import (
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/models"
	"github.com/tomp332/bruteForcer/src/utils"
)

type ICredentialsCrud struct{}

var CredentialsCrud = &ICredentialsCrud{}

// Get gets the credentials with the given id
func (c *ICredentialsCrud) Get(limit, offset uint) ([]models.IReadCredentials, error) {
	var objSlice []*models.CredentialsDTO
	err := src.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&objSlice).Error
	if err != nil {
		return nil, err
	}
	result := make([]models.IReadCredentials, len(objSlice))
	err = utils.CopyStructFields(objSlice, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Add adds the given credentials to the database
func (c *ICredentialsCrud) Add(creds []*models.ICredentialsCreate) ([]models.IReadCredentials, error) {
	credsModelSlice := make([]models.CredentialsDTO, len(creds))
	for i, credsBase := range creds {
		credsModelSlice[i] = models.CredentialsDTO{
			ICredentialsCreate: *credsBase,
			CustomORMModel:     models.CustomORMModel{},
		}
	}
	result := src.MainDB.Create(credsModelSlice)
	if result.Error != nil {
		return nil, result.Error
	}
	addedCredentials := make([]models.IReadCredentials, len(credsModelSlice))
	err := utils.CopyStructFields(addedCredentials, &creds)
	if err != nil {
		return nil, err
	}
	return addedCredentials, nil
}

// Update updates the credentials with the given id
func (c *ICredentialsCrud) Update(creds []*models.IUpdateCredentials) ([]models.CredentialsDTO, error) {
	updatedCredentials := make([]models.CredentialsDTO, len(creds))
	for i, updateSchema := range creds {
		updatedCredentials[i] = models.CredentialsDTO{
			ICredentialsCreate: models.ICredentialsCreate{
				Username: updateSchema.Username,
				Password: updateSchema.Password,
			},
			CustomORMModel: models.CustomORMModel{
				ID: updateSchema.ID,
			},
		}
	}
	result := src.MainDB.Save(updatedCredentials)
	if result.Error != nil {
		return nil, result.Error
	}
	return updatedCredentials, nil
}

// Delete deletes the credentials with the given id
func (c *ICredentialsCrud) Delete(id uint) error {
	result := src.MainDB.Delete(&models.CredentialsDTO{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
