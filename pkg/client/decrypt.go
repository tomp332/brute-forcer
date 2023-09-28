package client

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"github.com/tomp332/gobrute/pkg/client/plugins"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"io"
	"os"
	"time"
)

func DecryptWrapper(decryptionTask *internalTypes.DecryptionTask, numWorkers int) error {
	file, err := os.Open(decryptionTask.WordlistPath)
	if err != nil {
		logrus.Fatal("Error opening file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logrus.Fatal("Error closing file:", err)
		}
	}(file)
	currentPlugin := plugins.GetPlugin(decryptionTask.Mode)

	// Channels
	encryptJobChannel := make(chan internalTypes.EncryptionTask)
	wordlistChannel := make(chan string)
	encResultChannel := make(chan internalTypes.EncryptionTaskResult)
	notifyChannel := make(chan struct{}) // flag for decryption success

	go readWordlist(wordlistChannel, file, notifyChannel)
	for i := 0; i < numWorkers; i++ {
		go EncryptionWorker(currentPlugin, encryptJobChannel, encResultChannel)
	}
	go passwordReadCallback(decryptionTask, wordlistChannel, encryptJobChannel)

	return resultsChannelCallback(encResultChannel, decryptionTask)
}

func resultsChannelCallback(resultsChannel chan internalTypes.EncryptionTaskResult, decryptionTask *internalTypes.DecryptionTask) error {
	// Process the decrypted lines received from the output channel
	for hashResult := range resultsChannel {
		if hashResult.EncryptedHash == decryptionTask.TargetHash {
			decryptionTask.EndTime = time.Since(decryptionTask.StartTime)
			decryptionTask.ResultPassword = hashResult.TargetPassword
			logrus.WithFields(logrus.Fields{"passwordAttempts": decryptionTask.NumAttempts,
				"elapsedTime": decryptionTask.EndTime.Seconds(), "mode": decryptionTask.Mode,
				"modeName": plugins.GetPlugin(decryptionTask.Mode),
				"hash":     decryptionTask.TargetHash, "password": decryptionTask.ResultPassword}).Info("Operation successful")
			return nil
		}
	}
	decryptionTask.EndTime = time.Since(decryptionTask.StartTime)
	logrus.WithFields(logrus.Fields{"elapsedTime": decryptionTask.EndTime.Seconds(),
		"passwordAttempts": decryptionTask.NumAttempts, "mode": decryptionTask.Mode,
		"hash": decryptionTask.TargetHash}).Error("Operation failed, could not find a proper password to decrypt given hash.")
	return nil
}

func passwordReadCallback(decryptionTask *internalTypes.DecryptionTask, inputChannel chan string, outputChannel chan internalTypes.EncryptionTask) {
	for password := range inputChannel {
		decryptionTask.NumAttempts++ // Increment the number of attempts
		outputChannel <- internalTypes.EncryptionTask{
			TargetPassword: password,
			ResultHash:     decryptionTask.TargetHash,
		}
	}
}

func readWordlist(inputChannel chan string, file io.Reader, notifyChannel chan struct{}) {
	defer close(inputChannel)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputChannel <- line
		select {
		case <-notifyChannel:
			break
		default:
		}
	}
}
