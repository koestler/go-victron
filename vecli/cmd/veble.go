package cmd

import (
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/tinygoble"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

// vebleCmd defines the command to access BLE devices
var vebleCmd = &cobra.Command{
	Use:   "veble",
	Short: "Interact with Victron devices over bluetooth.",
	Long:  `Can list Victron received devices and decrypt messages given the mac and decryption key.`,
}

func init() {
	vebleCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(vebleCmd)
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for Victron devices.",
	Long:  `Scan for Victron devices and list them using the default bluetooth adapter.`,
	Run:   runScan,
}

func runScan(cmd *cobra.Command, args []string) {
	l := log.DefaultLogger{}

	a, err := tinygoble.NewDefaultAdapter(l)
	if err != nil {
		l.Printf("error creating adapter: %s", err)
		os.Exit(2)
	}
	a.RegisterDefaultListener(func(mac string, rssi int, localName string) {
		l.Printf("discovered : %s %d %s", mac, rssi, localName)
	})
	a.Run()
	defer a.Close()

	l.Printf("Scanning for Victron devices. press ctrl+c to abort...")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
