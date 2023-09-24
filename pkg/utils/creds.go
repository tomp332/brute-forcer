package utils

import (
	"github.com/tomp332/gobrute/pkg/types"
)

// TransformDTOCredentials converts a slice of CredentialsDTO to a slice of IReadCredentials
func TransformDTOCredentials(credsDtoSlice *[]types.CredentialsDTO) []types.IReadCredentials {
	readCredentials := make([]types.IReadCredentials, len(*credsDtoSlice))
	for i, obj := range *credsDtoSlice {
		readCredentials[i] = types.IReadCredentials{
			CredentialsBase: obj.CredentialsBase,
			ID:              obj.CustomORMModel.ID,
			CreatedAt:       obj.CustomORMModel.CreatedAt,
		}
	}
	return readCredentials
}
