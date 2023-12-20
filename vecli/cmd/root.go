package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vecli",
	Short: "Simple CLI for the go-victron library",
	Long: `This is a simple command line interface to connect to victron devices over VE.Direct and BLE.
It is mainly intended to be used for debugging and testing purposes.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable debug log output.")
}
