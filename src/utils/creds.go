package utils

import (
	"github.com/tomp332/bruteForcer/src/models"
)

// TransformDTOCredentials converts a slice of CredentialsDTO to a slice of IReadCredentials
func TransformDTOCredentials(credsDtoSlice *[]models.CredentialsDTO) []models.IReadCredentials {
	readCredentials := make([]models.IReadCredentials, len(*credsDtoSlice))
	for i, obj := range *credsDtoSlice {
		readCredentials[i] = models.IReadCredentials{
			CredentialsBase: obj.CredentialsBase,
			ID:              obj.CustomORMModel.ID,
			CreatedAt:       obj.CustomORMModel.CreatedAt,
		}
	}
	return readCredentials
}
