package crud

import (
	"github.com/tomp332/gobrute/pkg"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"github.com/tomp332/gobrute/pkg/utils"
	"log"
)

type ICredentialsCrud struct{}

var CredentialsCrud = &ICredentialsCrud{}

// Get gets the credentials with the given id
func (c *ICredentialsCrud) Get(limit, offset uint) ([]internalTypes.IReadCredentials, error) {
	var fetchedCreds []internalTypes.CredentialsDTO
	err := pkg.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&fetchedCreds).Error
	if err != nil {
		return nil, err
	}
	return utils.TransformDTOCredentials(&fetchedCreds), nil

}

// Add adds the given credentials to the database
func (c *ICredentialsCrud) Add(creds []internalTypes.ICredentialsCreate) ([]internalTypes.IReadCredentials, error) {
	credsModelSlice := make([]internalTypes.CredentialsDTO, len(creds))
	for i, credsBase := range creds {
		err := utils.CopyStructFields(credsBase, &credsModelSlice[i])
		if err != nil {
			return nil, err
		}
	}
	result := pkg.MainDB.Create(credsModelSlice)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Added new credentials to DB, count: %d", len(credsModelSlice))
	return utils.TransformDTOCredentials(&credsModelSlice), nil
}

// Update updates the credentials with the given id
func (c *ICredentialsCrud) Update(creds []*internalTypes.IUpdateCredentials) ([]internalTypes.IReadCredentials, error) {
	updatedCredentials := make([]internalTypes.CredentialsDTO, len(creds))
	for i, updateSchema := range creds {
		updatedCredentials[i] = internalTypes.CredentialsDTO{
			CustomORMModel: internalTypes.CustomORMModel{
				ID: updateSchema.ID,
			},
			CredentialsBase: updateSchema.CredentialsBase,
		}
	}
	result := pkg.MainDB.Save(updatedCredentials)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Updated credentials information with id %d", updatedCredentials[0].ID)
	return utils.TransformDTOCredentials(&updatedCredentials), nil
}

// Delete deletes the credentials with the given id
func (c *ICredentialsCrud) Delete(id uint) error {
	result := pkg.MainDB.Delete(&internalTypes.CredentialsDTO{}, id)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Deleted credentials with id %d", id)
	return nil
}
