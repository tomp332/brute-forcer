package crud

import (
	"github.com/tomp332/gobrute/src"
	"github.com/tomp332/gobrute/src/models"
	"github.com/tomp332/gobrute/src/plugins"
	"github.com/tomp332/gobrute/src/utils"
	"log"
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
	go func() {
		_, err := ExecuteBrute(bruteForceTasks)
		if err != nil {
			log.Fatalf("Error starting brute force action for provided hashes")
		}
	}()
	return utils.TransformDTOBruteForce(&bruteTasksSlice), nil
}

func ExecuteBrute(bruteForceTasks []models.IBruteForceCreate) (string, error) {
	// Create a map to store the split slices.
	splitTasksByAlgo := make(map[string][]models.IBruteForceCreate)
	// Iterate through the sorted slice and split by "Algorithm".
	for _, task := range bruteForceTasks {
		splitTasksByAlgo[task.Algorithm] = append(splitTasksByAlgo[task.Algorithm], task)
	}
	for algorithm, tasks := range splitTasksByAlgo {
		log.Printf("Handeling tasks of type %s", algorithm)
		obj := plugins.EncodingPluginsMap[algorithm]
		for _, bruteTask := range tasks {
			bruteTask := bruteTask
			go func() {
				decoded, err := obj.Decode(bruteTask.Hash)
				log.Printf("Hash: %s, Decoding: %s", bruteTask.Hash, decoded)
				if err != nil {
					log.Printf("Error decoding hash: %s. Error: %s", bruteTask.Hash, err)
				} else {
					// Add decoded hash to DB

				}
			}()
		}
	}
	return "", nil
}
