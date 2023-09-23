package utils

import "github.com/tomp332/gobrute/src/models"

// TransformDTOBruteForce converts a slice of BruteForceDTO to a slice of IBruteForceRead
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
