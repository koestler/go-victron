package cmd

import (
	"context"
	"fmt"
	"github.com/koestler/go-victron/vedirect"
	"github.com/koestler/go-victron/vedirectapi"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// vedirectCmd represents the vedirect command
var vedirectCmd = &cobra.Command{
	Use:   "vedirect",
	Short: "Connect to VE.Direct device.",
	Long: `Connect to VE.Direct device, sends a ping, gets the device id and uses it to get the product type.
Then it uses the register list to fetch all known registers and outputs it's values.`,
	Run: runVedirect,
}

func init() {
	rootCmd.AddCommand(vedirectCmd)
	vedirectCmd.Flags().StringP("device", "d", "", "Serial device to connect to. E.g. /dev/ttyUSB0")
	_ = vedirectCmd.MarkFlagRequired("device")

	vedirectCmd.Flags().StringP(
		"io-log", "i", "",
		"File to write an input/output log to. This is useful to create unit tests.",
	)
}

func runVedirect(cmd *cobra.Command, args []string) {
	var vdConfig vedirect.Config

	if verbose, err := cmd.Flags().GetBool("verbose"); err == nil && verbose {
		vdConfig.DebugLogger = log.Default()
	}

	if ioLog, err := cmd.Flags().GetString("io-log"); err == nil && ioLog != "" {
		fl, err := vedirectapi.NewFileLogger(ioLog)
		if err != nil {
			fmt.Printf("error creating io log file: %s\n", err)
			return
		}
		defer fl.Close()
		vdConfig.IoLogger = fl
	}

	time0 := time.Now()
	api, err := vedirectapi.NewSerialRegisterApi(cmd.Flag("device").Value.String(), vdConfig)
	if err != nil {
		fmt.Printf("error creating api: %s\n", err)
		return
	}
	defer api.Close()

	time1 := time.Now()

	var values vedirectapi.RegisterValues
	if rv, err := api.ReadAllRegisters(context.Background()); err != nil {
		fmt.Printf("error fetching registers: %s\n", err)
	} else {
		values = rv
	}

	time2 := time.Now()

	list := values.GetList()

	fmt.Printf("fetched %d registers, initialization took: %s, fetching took: %s\n",
		len(list),
		time1.Sub(time0),
		time2.Sub(time1),
	)

	// output list
	for _, l := range list {
		fmt.Println(l)
	}
}
