package commands

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tomp332/gobrute/pkg/client"
	"github.com/tomp332/gobrute/pkg/internalTypes"
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
	Long:    `Decrypt a hash using client`,
	Args:    cobra.ExactArgs(1),
	Example: `client decrypt -m [mode] -w [wordlist] [secret hash]`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateWordlistFlags()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if decryptionMode == 0 {
			return errors.New("the decryption mode flag is required")
		}
		decryptionTask := &internalTypes.DecryptionTask{
			TargetHash:   args[0],
			Mode:         decryptionMode,
			WordlistPath: wordlistPath,
			NumAttempts:  0,
			Task: internalTypes.Task{
				StartTime: time.Now(),
			},
		}
		err := client.DecryptWrapper(decryptionTask, numWorkers)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	DecryptCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "client [options] -v")
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
