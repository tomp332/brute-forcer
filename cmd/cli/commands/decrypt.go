package commands

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tomp332/gobrute/pkg/cli"
	"github.com/tomp332/gobrute/pkg/cli/plugins"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"io"
	"os"
	"time"
)

var (
	numWorkers     int
	decryptionMode int16
	verboseFlag    bool
	wordlistSlice  *[]string
	wordlistPath   string
)
var DecryptCmd = &cobra.Command{
	Use:     "decrypt",
	Short:   "Decrypt a hash",
	Long:    `Decrypt a hash using gobrute`,
	Args:    cobra.ExactArgs(1),
	Example: `gobrute decrypt -m [mode] -w [wordlist] [secret hash]`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateWordlistFlags()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if decryptionMode == 0 {
			return errors.New("the decryption mode flag is required")
		}
		decryptionTask := &internalTypes.DecryptionTask{
			Hash:         args[0],
			Mode:         decryptionMode,
			WordlistPath: wordlistPath,
			NumAttempts:  0,
			Task: internalTypes.Task{
				StartTime: time.Now(),
			},
		}
		err := DecryptWrapper(decryptionTask)
		if err != nil {
			return err
		}
		return nil
	},
}

func DecryptWrapper(decryptionTask *internalTypes.DecryptionTask) error {
	file, err := os.Open(decryptionTask.WordlistPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error closing file:", err)
		}
	}(file)
	currentPlugin := plugins.GetPlugin(decryptionTask.Mode)

	// Channels
	encryptJobChannel := make(chan internalTypes.EncryptionTask)
	wordlistChannel := make(chan string)
	encResultChannel := make(chan internalTypes.TaskResult)
	notifyChannel := make(chan struct{}) // flag for decryption success

	go readWordlist(wordlistChannel, file, notifyChannel)
	for i := 0; i < numWorkers; i++ {
		go cli.EncryptionWorker(currentPlugin, encryptJobChannel, encResultChannel)
	}
	go handleWordReadCallback(decryptionTask, wordlistChannel, encryptJobChannel)

	return resultsChannelCallback(encResultChannel, decryptionTask)
}

func resultsChannelCallback(resultsChannel chan internalTypes.TaskResult, decryptionTask *internalTypes.DecryptionTask) error {
	// Process the decrypted lines received from the output channel
	for hashResult := range resultsChannel {
		if hashResult.Hash == decryptionTask.Hash {
			decryptionTask.EndTime = time.Since(decryptionTask.StartTime)
			decryptionTask.Password = hashResult.Password
			log.WithFields(log.Fields{"passwordAttempts": decryptionTask.NumAttempts, "elapsedTime": decryptionTask.EndTime.Seconds(), "mode": decryptionTask.Mode,
				"hash": decryptionTask.Hash, "password": decryptionTask.Password}).Info("Operation successful")
			return nil
		}
	}
	decryptionTask.EndTime = time.Since(decryptionTask.StartTime)
	log.WithFields(log.Fields{"elapsedTime": decryptionTask.EndTime.Seconds(), "passwordAttempts": decryptionTask.NumAttempts, "mode": decryptionTask.Mode,
		"hash": decryptionTask.Hash}).Error("Operation failed, could not find a proper password to decrypt given hash.")
	return nil
}

func handleWordReadCallback(decryptionTask *internalTypes.DecryptionTask, inputChannel chan string, outputChannel chan internalTypes.EncryptionTask) {
	for password := range inputChannel {
		decryptionTask.NumAttempts++ // Increment the number of attempts
		outputChannel <- internalTypes.EncryptionTask{
			Password: password,
			Hash:     decryptionTask.Hash,
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

func init() {
	DecryptCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "gobrute [options] -v")
	DecryptCmd.Flags().Int16VarP(&decryptionMode, "mode", "m", 0, "-m [mode]")
	DecryptCmd.Flags().StringVarP(&wordlistPath, "wordlist-file", "f", "",
		"-w [wordlist full file path]")
	DecryptCmd.Flags().IntVarP(&numWorkers, "workers", "w", 1, "-w [num of workers]")
	wordlistSlice = DecryptCmd.Flags().StringSliceP("wordlist-array", "l", []string{},
		"-l [a,b,c...]")
	DecryptCmd.MarkFlagsMutuallyExclusive("wordlist-file", "wordlist-array")
	_ = DecryptCmd.MarkFlagRequired("mode")
}

func validateWordlistFlags() error {
	if verboseFlag {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	//Validate that at least one of the wordlist flags is used
	if wordlistPath == "" && len(*wordlistSlice) == 0 {
		return errors.New("one of the wordlist flags must be used")
	}
	// Validate that only one of the wordlist flags is used
	if wordlistPath != "" && len(*wordlistSlice) > 0 {
		return errors.New("only one of the wordlist flags can be used")
	}
	// Validate that the wordlist file exists
	if wordlistPath != "" {
		if _, err := os.Stat(wordlistPath); os.IsNotExist(err) {
			return errors.New("the wordlist file does not exist")
		}
	}
	return nil
}
