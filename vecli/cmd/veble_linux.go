package cmd

import (
	"bytes"
	"fmt"
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/tinygoble"
	"github.com/koestler/go-victron/veble"
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
	a.RegisterDefaultListener(func(mac veble.MAC, rssi int, localName string) {
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
	seen := make(map[veble.MAC]struct{})
	a.RegisterDefaultListener(func(mac veble.MAC, rssi int, localName string) {
		if _, ok := seen[mac]; ok {
			return
		}
		seen[mac] = struct{}{}
		fmt.Printf("discovered : mac=%s, RSSI=%d, name=%s\n", mac, rssi, localName)
	})
	defer a.Close()

	for _, arg := range args {
		p, err := parseMacKeyPair(arg)
		if err != nil {
			fmt.Println("error parsing mac key pair:", err)
			os.Exit(2)
		}

		a.RegisterMacListener(p.mac, macListener(p, debugLogger))
	}

	fmt.Println("Showing received packets. Press ctrl+c to abort...")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

func macListener(pair macKeyPair, l log.Logger) tinygoble.MacListener {
	var lastData []byte
	return func(rssi int, localName string, victronData []byte) {
		// ignore duplicate packets
		if bytes.Equal(lastData, victronData) {
			return
		}
		lastData = victronData

		ef, err := veble.DecodeFrame(victronData, l)
		if err != nil {
			fmt.Println("error decoding frame:", err)
			return
		}

		fmt.Printf("received packet MAC=%s, RSSI=%d, name=%s, IV=%d\n", pair.mac, rssi, localName, ef.IV)

		df, err := veble.DecryptFrame(ef, pair.key, l)
		if err != nil {
			fmt.Println("error decrypting frame:", err)
			return
		}

		regs, err := veble.Decode(df.RecordType, df.DecryptedBytes)
		if err != nil {
			fmt.Println("error decoding registers:", err)
			return
		}

		for _, nr := range regs.NumberRegisters() {
			fmt.Println("- ", nr)
		}
		for _, er := range regs.EnumRegisters() {
			fmt.Println("- ", er)
		}
		for _, flr := range regs.FieldListRegisters() {
			fmt.Println("- ", flr)
		}
	}
}
