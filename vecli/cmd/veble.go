package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/vebleapi"
	"github.com/koestler/go-victron/veblerecord"
	"github.com/koestler/go-victron/veconst"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
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

	a, err := vebleapi.NewDefaultAdapter(veconst.BleManufacturerId, debugLogger)
	if err != nil {
		fmt.Println("error creating adapter:", err)
		os.Exit(2)
	}
	defer a.Close()

	defaultListener := a.RegisterDefaultListener()
	defer defaultListener.End()
	go func() {
		for p := range defaultListener.Drain() {
			fmt.Printf("discovered : address=%s, RSSI=%d, name=%s\n", p.Address(), p.RSSI(), p.Name())
		}
	}()

	fmt.Println("Scanning for Victron devices. Press ctrl+c to abort...")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

func runDecode(cmd *cobra.Command, args []string) {
	debugLogger := getLogger(cmd)
	printRegister := false
	if v, err := cmd.Flags().GetBool("print-registers"); err == nil && v {
		printRegister = true
	}

	a, err := vebleapi.NewDefaultAdapter(veconst.BleManufacturerId, debugLogger)
	if err != nil {
		fmt.Println("error creating adapter:", err)
		os.Exit(2)
	}
	defer a.Close()

	seen := make(map[string]struct{})
	defaultListener := a.RegisterDefaultListener()
	defer defaultListener.End()
	go func() {
		for p := range defaultListener.Drain() {
			if _, ok := seen[p.Address()]; ok {
				continue
			}
			seen[p.Address()] = struct{}{}
			fmt.Printf("discovered : address=%s, RSSI=%d, name=%s\n", p.Address(), p.RSSI(), p.Name())
		}
	}()

	ctx, cancel := context.WithCancel(cmd.Context())

	for _, arg := range args {
		p, err := parseMacKeyPair(arg)
		if err != nil {
			fmt.Println("error parsing mac key pair:", err)
			os.Exit(2)
		}

		api := vebleapi.NewApi(a, p.name, p.key, debugLogger)

		go api.StreamRegisters(ctx, func(rssi int, registers veblerecord.Registers) {
			fmt.Printf("%s sent with RSSI=%d:\n", p.name, rssi)

			if printRegister {
				for _, r := range registers.NumberRegisters() {
					fmt.Printf("- %s\n", r)
				}
				for _, r := range registers.EnumRegisters() {
					fmt.Printf("- %s\n", r)
				}
				for _, r := range registers.FieldListRegisters() {
					fmt.Printf("- %s\n", r)
				}
			} else {
				j, err := json.MarshalIndent(registers, "", "  ")
				if err != nil {
					fmt.Println("error encoding record:", err)
					return
				}
				fmt.Println(string(j))
			}
		})
	}

	fmt.Println("Showing received packets. Press ctrl+c to abort...")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	cancel()
}
