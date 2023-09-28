package cli

import (
	log "github.com/sirupsen/logrus"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

func EncryptionWorker(plugin internalTypes.GoBrutePlugin, encryptionTasks <-chan internalTypes.EncryptionTask, results chan<- internalTypes.TaskResult) {
	for encryptionTask := range encryptionTasks {
		log.WithFields(log.Fields{"password": encryptionTask.Password, "hash": encryptionTask.Hash}).Debug("Started evaluating password on worker")
		result := internalTypes.TaskResult{
			Password: encryptionTask.Password,
		}
		err := plugin.Execute(&result)
		if err != nil {
			log.WithFields(log.Fields{"password": encryptionTask.Password}).Error("Failed to evaluate password")
			continue
		}
		log.WithFields(log.Fields{"password": encryptionTask.Password, "hash": encryptionTask.Hash}).Debug("Finished evaluating password on worker")
		results <- result
	}
}
