package commands

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tomp332/gobrute/pkg/cli/plugins"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"os"
	"sync"
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
		task := &internalTypes.Task{
			Hash:         args[0],
			Mode:         decryptionMode,
			WordlistPath: wordlistPath,
		}
		err := DecryptWrapper(task)
		if err != nil {
			return err
		}
		return nil
	},
}

func DecryptWrapper(t *internalTypes.Task) error {
	file, err := os.Open(t.WordlistPath)
	startTime := time.Now()
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error closing file:", err)
		}
	}(file)

	// Create channels for passing lines to be encrypted and receiving results
	inputChannel := make(chan string)
	outputChannel := make(chan *internalTypes.PluginResult)
	done := make(chan struct{})
	var wg sync.WaitGroup

	// Start worker goroutines
	currentPlugin := plugins.GetPlugin(t.Mode)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range inputChannel {
				result := &internalTypes.PluginResult{
					Password: line,
				}
				log.WithFields(log.Fields{"target": t.Hash, "password": line}).Debug("Executing decryption for password")
				err := currentPlugin.Execute(result)
				if err != nil {
					log.WithFields(log.Fields{"password": line}).Error("Failed to execute function on password")
					continue
				}
				outputChannel <- result
			}
		}()
	}

	// Read the file line by line and send lines to be decrypted
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			inputChannel <- line
			select {
			case <-done:
				// If any worker succeeded, stop sending lines
				break
			default:
			}
		}
		close(inputChannel)
	}()

	// Close the output channel when all workers are done
	go func() {
		wg.Wait()
		close(outputChannel)
	}()

	// Process the decrypted lines received from the output channel
	for hashResult := range outputChannel {
		if hashResult.Hash == t.Hash {
			elapsed := time.Since(startTime)
			t.PlaintText = hashResult.Password
			log.WithFields(log.Fields{"elapsedTime": elapsed, "mode": t.Mode, "hash": t.Hash, "password": t.PlaintText}).Info("Operation successful")
			return nil
		}
	}
	elapsed := time.Since(startTime)
	log.WithFields(log.Fields{"elapsedTime": elapsed, "mode": t.Mode, "hash": t.Hash}).Error("Operation failed, could not find a proper password to decrypt given hash.")
	return nil
}

func init() {
	DecryptCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "gobrute [options] -v")
	DecryptCmd.Flags().Int16VarP(&decryptionMode, "mode", "m", 0, "-m [mode]")
	DecryptCmd.Flags().StringVarP(&wordlistPath, "wordlist-file", "f", "", "-w [wordlist full file path]")
	DecryptCmd.Flags().IntVarP(&numWorkers, "workers", "w", 1, "-w [num of workers]")
	wordlistSlice = DecryptCmd.Flags().StringSliceP("wordlist-array", "l", []string{}, "-l [a,b,c...]")
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
