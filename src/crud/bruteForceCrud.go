package crud

import (
	"log"

	"github.com/tomp332/gobrute/src"
	"github.com/tomp332/gobrute/src/models"
	"github.com/tomp332/gobrute/src/utils"
)

type IBruteForceCrud struct{}

var BruteForceCrud = &IBruteForceCrud{}

// Get gets the credentials with the given id
func (c *IBruteForceCrud) Get(limit, offset uint) ([]models.IBruteForceRead, error) {
	var fetchedTask []models.BruteForceDTO
	err := src.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&fetchedTask).Error
	if err != nil {
		return nil, err
	}
	return utils.TransformDTOBruteForce(&fetchedTask), nil
}

// Add adds the given credentials to the database
func (c *IBruteForceCrud) Add(bruteForceTasks []models.IBruteForceCreate) ([]models.IBruteForceRead, error) {
	bruteTasksSlice := make([]models.BruteForceDTO, len(bruteForceTasks))
	for i, bruteTask := range bruteForceTasks {
		err := utils.CopyStructFields(bruteTask, &bruteTasksSlice[i])
		if err != nil {
			return nil, err
		}
	}
	result := src.MainDB.Create(bruteTasksSlice)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("Added new brute force tasks to DB, count: %d", len(bruteTasksSlice))
	return utils.TransformDTOBruteForce(&bruteTasksSlice), nil
}
