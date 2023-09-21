package utils

import "github.com/tomp332/gobrute/src/models"

// TransformDTOCredentials converts a slice of CredentialsDTO to a slice of IReadCredentials
func TransformDTOBruteForce(bruteTasksToSlice *[]models.BruteForceDTO) []models.IBruteForceRead {
	readBruteTasks := make([]models.IBruteForceRead, len(*bruteTasksToSlice))
	for i, obj := range *bruteTasksToSlice {
		readBruteTasks[i] = models.IBruteForceRead{
			BruteForceBase: obj.BruteForceBase,
			ID:             obj.CustomORMModel.ID,
			CreatedAt:      obj.CustomORMModel.CreatedAt,
		}
	}
	return readBruteTasks
}
