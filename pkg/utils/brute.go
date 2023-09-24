package utils

import (
	"github.com/tomp332/gobrute/pkg/types"
)

// TransformDTOBruteForce converts a slice of BruteForceDTO to a slice of IBruteForceRead
func TransformDTOBruteForce(bruteTasksToSlice *[]types.BruteForceDTO) []types.IBruteForceRead {
	readBruteTasks := make([]types.IBruteForceRead, len(*bruteTasksToSlice))
	for i, obj := range *bruteTasksToSlice {
		readBruteTasks[i] = types.IBruteForceRead{
			BruteForceBase: obj.BruteForceBase,
			ID:             obj.CustomORMModel.ID,
			CreatedAt:      obj.CustomORMModel.CreatedAt,
		}
	}
	return readBruteTasks
}
