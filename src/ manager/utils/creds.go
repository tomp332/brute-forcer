package utils

import "github.com/tomp332/gobrute/src/ manager/managerTypes"

// TransformDTOCredentials converts a slice of CredentialsDTO to a slice of IReadCredentials
func TransformDTOCredentials(credsDtoSlice *[]managerTypes.CredentialsDTO) []managerTypes.IReadCredentials {
	readCredentials := make([]managerTypes.IReadCredentials, len(*credsDtoSlice))
	for i, obj := range *credsDtoSlice {
		readCredentials[i] = managerTypes.IReadCredentials{
			CredentialsBase: obj.CredentialsBase,
			ID:              obj.CustomORMModel.ID,
			CreatedAt:       obj.CustomORMModel.CreatedAt,
		}
	}
	return readCredentials
}
