package cmd

import (
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/tinygoble"
	"github.com/koestler/go-victron/veble"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func runScan(_ *cobra.Command, _ []string) {
	l := log.DefaultLogger{}

	a, err := tinygoble.NewDefaultAdapter(l)
	if err != nil {
		l.Printf("error creating adapter: %s", err)
		os.Exit(2)
	}
	a.RegisterDefaultListener(func(mac veble.MAC, rssi int, localName string) {
		l.Printf("discovered : mac=%s, RSSI=%d, name=%s", mac, rssi, localName)
	})
	defer a.Close()

	l.Printf("Scanning for Victron devices. press ctrl+c to abort...")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

func runDecode(cmd *cobra.Command, args []string) {

}
