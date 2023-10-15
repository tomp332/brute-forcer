package client

import (
	log "github.com/sirupsen/logrus"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

func EncryptionWorker(plugin internalTypes.GoBrutePlugin, encryptionTasks <-chan internalTypes.EncryptionTask, results chan<- internalTypes.EncryptionTaskResult) {
	for encryptionTask := range encryptionTasks {
		log.WithFields(log.Fields{"password": encryptionTask.TargetPassword, "targetHash": encryptionTask.ResultHash}).Debug("Started evaluating password on worker")
		// TODO: Parse CLI arguments and fill this struct
		result := internalTypes.EncryptionTaskResult{
			TargetPassword: encryptionTask.TargetPassword,
		}
		err := plugin.Execute(&result)
		if err != nil {
			log.WithFields(log.Fields{"password": encryptionTask.TargetPassword}).Error("Failed to evaluate password")
			continue
		}
		log.WithFields(log.Fields{"password": encryptionTask.TargetPassword, "evaluatedHash": result.EncryptedHash, "targetHash": encryptionTask.ResultHash}).Debug("Finished evaluating password on worker")
		results <- result
	}
}
