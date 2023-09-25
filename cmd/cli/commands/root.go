package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobrute-cli",
	Short: "gobrute-cli is a CLI tool for gobrute",
	Long:  `gobrute-cli is a CLI tool for gobrute. It is a tool for brute forcing`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
