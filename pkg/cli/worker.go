package cli

import (
	log "github.com/sirupsen/logrus"
	"github.com/tomp332/gobrute/pkg/internalTypes"
)

func EncryptionWorker(encryptionTasks <-chan internalTypes.PluginResult, plugin internalTypes.GoBrutePlugin, results chan<- internalTypes.PluginResult) {
	for task := range encryptionTasks {
		log.WithFields(log.Fields{"password": task.Password}).Debug("Started evaluating password..")
		result := internalTypes.PluginResult{
			Password: task.Password,
		}
		err := plugin.Execute(&result)
		if err != nil {
			log.WithFields(log.Fields{"password": task.Password}).Error("Failed to evaluate password")
			continue
		}
		log.WithFields(log.Fields{"password": task.Password}).Debug("Finished evaluating password..")
		results <- result
	}
}
