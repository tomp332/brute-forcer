package crud

import (
	"github.com/tomp332/gobrute/cmd/manager/managerTypes"
	"github.com/tomp332/gobrute/cmd/manager/utils"
	"log"
)

type ICredentialsCrud struct{}

var CredentialsCrud = &ICredentialsCrud{}

// Get gets the credentials with the given id
func (c *ICredentialsCrud) Get(limit, offset uint) ([]managerTypes.IReadCredentials, error) {
	var fetchedCreds []managerTypes.CredentialsDTO
	err := manager.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&fetchedCreds).Error
	if err != nil {
		return nil, err
	}
	return utils.TransformDTOCredentials(&fetchedCreds), nil

}

// Add adds the given credentials to the database
func (c *ICredentialsCrud) Add(creds []managerTypes.ICredentialsCreate) ([]managerTypes.IReadCredentials, error) {
	credsModelSlice := make([]managerTypes.CredentialsDTO, len(creds))
	for i, credsBase := range creds {
		err := utils.CopyStructFields(credsBase, &credsModelSlice[i])
		if err != nil {
			return nil, err
		}
	}
	result := manager.MainDB.Create(credsModelSlice)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Added new credentials to DB, count: %d", len(credsModelSlice))
	return utils.TransformDTOCredentials(&credsModelSlice), nil
}

// Update updates the credentials with the given id
func (c *ICredentialsCrud) Update(creds []*managerTypes.IUpdateCredentials) ([]managerTypes.IReadCredentials, error) {
	updatedCredentials := make([]managerTypes.CredentialsDTO, len(creds))
	for i, updateSchema := range creds {
		updatedCredentials[i] = managerTypes.CredentialsDTO{
			CustomORMModel: managerTypes.CustomORMModel{
				ID: updateSchema.ID,
			},
			CredentialsBase: updateSchema.CredentialsBase,
		}
	}
	result := manager.MainDB.Save(updatedCredentials)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Updated credentials information with id %d", updatedCredentials[0].ID)
	return utils.TransformDTOCredentials(&updatedCredentials), nil
}

// Delete deletes the credentials with the given id
func (c *ICredentialsCrud) Delete(id uint) error {
	result := manager.MainDB.Delete(&managerTypes.CredentialsDTO{}, id)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Deleted credentials with id %d", id)
	return nil
}
