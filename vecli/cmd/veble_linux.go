package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/mac"
	"github.com/koestler/go-victron/tinygoble"
	"github.com/koestler/go-victron/vebleapi"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func getLogger(cmd *cobra.Command) log.Logger {
	if verbose, err := cmd.Flags().GetBool("verbose"); err == nil && verbose {
		return log.DefaultLogger{
			Prefix: "ble: ",
		}
	}

	return log.NoOppLogger{}
}

func runScan(cmd *cobra.Command, _ []string) {
	debugLogger := getLogger(cmd)

	a, err := tinygoble.NewDefaultAdapter(debugLogger)
	if err != nil {
		fmt.Println("error creating adapter:", err)
		os.Exit(2)
	}
	a.RegisterDefaultListener(func(mac mac.MAC, rssi int, localName string) {
		fmt.Printf("discovered : mac=%s, RSSI=%d, name=%s\n", mac, rssi, localName)
	})
	defer a.Close()

	fmt.Println("Scanning for Victron devices. Press ctrl+c to abort...")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

func runDecode(cmd *cobra.Command, args []string) {
	debugLogger := getLogger(cmd)

	a, err := tinygoble.NewDefaultAdapter(debugLogger)
	if err != nil {
		fmt.Println("error creating adapter:", err)
		os.Exit(2)
	}
	seen := make(map[mac.MAC]struct{})
	a.RegisterDefaultListener(func(mac mac.MAC, rssi int, localName string) {
		if _, ok := seen[mac]; ok {
			return
		}
		seen[mac] = struct{}{}
		fmt.Printf("discovered : mac=%s, RSSI=%d, name=%s\n", mac, rssi, localName)
	})
	defer a.Close()

	ctx, cancel := context.WithCancel(cmd.Context())

	for _, arg := range args {
		p, err := parseMacKeyPair(arg)
		if err != nil {
			fmt.Println("error parsing mac key pair:", err)
			os.Exit(2)
		}

		api := vebleapi.NewRecordApi(a, p.mac, p.key, debugLogger)
		go api.Stream(ctx, func(rssi int, localName string, record any) {
			fmt.Printf("%s: %s sent with RSSI=%d:\n", p.mac, localName, rssi)
			j, err := json.MarshalIndent(record, "", "  ")
			if err != nil {
				fmt.Println("error encoding record:", err)
				return
			}
			fmt.Println(string(j))
		})
	}

	fmt.Println("Showing received packets. Press ctrl+c to abort...")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	cancel()
}
