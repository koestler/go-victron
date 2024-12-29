package cmd

import (
	"github.com/spf13/cobra"
)

// vebleCmd defines the command to access BLE devices
var vebleCmd = &cobra.Command{
	Use:   "veble",
	Short: "Interact with Victron devices over bluetooth.",
	Long:  "Can list Victron received devices and decrypt messages given the mac and decryption key.",
}

func init() {
	decodeCmd.Flags().BoolP("print-registers", "r", false, "Print registers instead of json.")

	vebleCmd.AddCommand(scanCmd)
	vebleCmd.AddCommand(decodeCmd)
	rootCmd.AddCommand(vebleCmd)
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for Victron devices.",
	Long:  "Scan for Victron devices and list them using the default bluetooth adapter.",
	Run:   runScan,
}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode all message from a specific Victron devices.",
	Long: `Decrypt and decode all messages from a specific Victron devices given their MAC address and encryption key.
Supply arguments in the form MAC=KEY as shown in the Victron Energy App. E.g. e675a31ea872=713f401f0b05beb18ec0937768162e4e`,
	Args: cobra.MinimumNArgs(1),
	Run:  runDecode,
}
