package utils

import (
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

// TransformDTOBruteForce converts a slice of BruteForceDTO to a slice of IBruteForceRead
func TransformDTOBruteForce(bruteTasksToSlice *[]internalTypes.BruteForceDTO) []internalTypes.IBruteForceRead {
	readBruteTasks := make([]internalTypes.IBruteForceRead, len(*bruteTasksToSlice))
	for i, obj := range *bruteTasksToSlice {
		readBruteTasks[i] = internalTypes.IBruteForceRead{
			BruteForceBase: obj.BruteForceBase,
			ID:             obj.CustomORMModel.ID,
			CreatedAt:      obj.CustomORMModel.CreatedAt,
		}
	}
	return readBruteTasks
}
