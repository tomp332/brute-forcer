package commands

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gobrute",
	Short: "A brief description of your application",
	Long:  `gobrute is a cli tool for a fast brute forcing attack.`,
}

func init() {
	rootCmd.AddCommand(DecryptCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
