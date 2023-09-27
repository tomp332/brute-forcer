package commands

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tomp332/gobrute/pkg/cli/plugins"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"os"
)

var decryptionMode int16
var wordlistPath string
var wordlistSlice *[]string
var verboseFlag bool
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
		encryptedHash, err := plugins.DecryptWrapper(task)
		if err != nil {
			return err
		}
		if len(encryptedHash) != 0 {
			log.WithFields(log.Fields{"mode": task.Mode, "hash": task.Hash, "password": task.PlaintText}).Info("Operation successful")
		}
		return nil
	},
}

func init() {
	DecryptCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", true, "gobrute [options] -v")
	DecryptCmd.Flags().Int16VarP(&decryptionMode, "mode", "m", 0, "-m [mode]")
	DecryptCmd.Flags().StringVarP(&wordlistPath, "wordlist", "w", "", "-w [wordlist full file path]")
	wordlistSlice = DecryptCmd.Flags().StringSliceP("wordlist-array", "l", []string{}, "-l [a,b,c...]")
	DecryptCmd.MarkFlagsMutuallyExclusive("wordlist", "wordlist-array")
	_ = DecryptCmd.MarkFlagRequired("mode")
	if verboseFlag {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func validateWordlistFlags() error {
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
