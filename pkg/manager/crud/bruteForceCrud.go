package crud

//
//import (
//	"github.com/tomp332/gobrute/pkg/client/plugins"
//	"github.com/tomp332/gobrute/pkg/internalTypes"
//	"github.com/tomp332/gobrute/pkg/manager"
//	utils2 "github.com/tomp332/gobrute/pkg/manager/utils"
//	"log"
//)
//
//type IBruteForceCrud struct{}
//
//var BruteForceCrud = &IBruteForceCrud{}
//
//// Get gets the credentials with the given id
//func (c *IBruteForceCrud) Get(limit, offset uint) ([]internalTypes.IBruteForceRead, error) {
//	var fetchedTask []internalTypes.BruteForceDTO
//	err := manager.MainDB.Scopes(NewPaginate(limit, offset).PaginatedResult).Find(&fetchedTask).Error
//	if err != nil {
//		return nil, err
//	}
//	return utils2.TransformDTOBruteForce(&fetchedTask), nil
//}
//
//// Add adds the given credentials to the database
//func (c *IBruteForceCrud) Add(bruteForceTasks []internalTypes.IBruteForceCreate) ([]internalTypes.IBruteForceRead, error) {
//	bruteTasksSlice := make([]internalTypes.BruteForceDTO, len(bruteForceTasks))
//	for i, bruteTask := range bruteForceTasks {
//		err := utils2.CopyStructFields(bruteTask, &bruteTasksSlice[i])
//		if err != nil {
//			return nil, err
//		}
//	}
//	result := manager.MainDB.Create(bruteTasksSlice)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	log.Printf("Added new brute force tasks to DB, count: %d", len(bruteTasksSlice))
//	go func() {
//		_, err := ExecuteBrute(bruteForceTasks)
//		if err != nil {
//			log.Fatalf("Error starting brute force action for provided hashes")
//		}
//	}()
//	return utils2.TransformDTOBruteForce(&bruteTasksSlice), nil
//}
//
//func ExecuteBrute(bruteForceTasks []internalTypes.IBruteForceCreate) (string, error) {
//	// Create a map to store the split slices.
//	splitTasksByAlgo := make(map[string][]internalTypes.IBruteForceCreate)
//	// Iterate through the sorted slice and split by "Algorithm".
//	for _, task := range bruteForceTasks {
//		splitTasksByAlgo[task.Algorithm] = append(splitTasksByAlgo[task.Algorithm], task)
//	}
//	for algorithm, tasks := range splitTasksByAlgo {
//		log.Printf("Handeling tasks of type %s", algorithm)
//		obj := plugins.GetPlugin(algorithm)
//		for _, bruteTask := range tasks {
//			bruteTask := bruteTask
//			go func() {
//				decoded, err := obj.Decode(bruteTask.Hash)
//				log.Printf("TargetHash: %s, Decoding: %s", bruteTask.Hash, decoded)
//				if err != nil {
//					log.Printf("Error decoding hash: %s. Error: %s", bruteTask.Hash, err)
//				} else {
//					// Add decoded hash to DB
//
//				}
//			}()
//		}
//	}
//	return "", nil
//}
