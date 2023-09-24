package utils

import "github.com/tomp332/gobrute/src/ manager/managerTypes"

// TransformDTOBruteForce converts a slice of BruteForceDTO to a slice of IBruteForceRead
func TransformDTOBruteForce(bruteTasksToSlice *[]managerTypes.BruteForceDTO) []managerTypes.IBruteForceRead {
	readBruteTasks := make([]managerTypes.IBruteForceRead, len(*bruteTasksToSlice))
	for i, obj := range *bruteTasksToSlice {
		readBruteTasks[i] = managerTypes.IBruteForceRead{
			BruteForceBase: obj.BruteForceBase,
			ID:             obj.CustomORMModel.ID,
			CreatedAt:      obj.CustomORMModel.CreatedAt,
		}
	}
	return readBruteTasks
}
