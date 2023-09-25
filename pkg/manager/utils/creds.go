package utils

import (
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

// TransformDTOCredentials converts a slice of CredentialsDTO to a slice of IReadCredentials
func TransformDTOCredentials(credsDtoSlice *[]internalTypes.CredentialsDTO) []internalTypes.IReadCredentials {
	readCredentials := make([]internalTypes.IReadCredentials, len(*credsDtoSlice))
	for i, obj := range *credsDtoSlice {
		readCredentials[i] = internalTypes.IReadCredentials{
			CredentialsBase: obj.CredentialsBase,
			ID:              obj.CustomORMModel.ID,
			CreatedAt:       obj.CustomORMModel.CreatedAt,
		}
	}
	return readCredentials
}
