package cmd

import (
	"fmt"
	"github.com/koestler/go-victron/vedirectapi"
	"github.com/spf13/cobra"
	"log"
	"sort"
	"time"
)

// vedirectCmd represents the vedirect command
var vedirectCmd = &cobra.Command{
	Use:   "vedirect",
	Short: "Connect to VE.Direct device.",
	Long: `Connect to VE.Direct device, sends a ping, gets the device id and uses it to get the product type.
Then it uses the register list to fetch all known registers and outputs it's values.`,
	Run: vedirect,
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

func vedirect(cmd *cobra.Command, args []string) {
	cfg := vedirectapi.Config{
		SerialDevice: cmd.Flag("device").Value.String(),
	}

	if verbose, err := cmd.Flags().GetBool("verbose"); err == nil && verbose {
		cfg.DebugLogger = log.Default()
	}

	if ioLog, err := cmd.Flags().GetString("io-log"); err == nil && ioLog != "" {
		fl, err := newFileLogger(ioLog)
		if err != nil {
			fmt.Printf("error creating io log file: %s\n", err)
			return
		}
		defer fl.Close()
		cfg.IoLogger = fl
	}

	time0 := time.Now()
	api, err := vedirectapi.NewApi(&cfg)
	if err != nil {
		fmt.Printf("error creating api: %s\n", err)
		return
	}
	defer api.Close()

	time1 := time.Now()

	var list []ls

	if rv, err := api.FetchAllRegisters(); err != nil {
		fmt.Printf("error fetching registers: %s\n", err)
	} else {
		// create a list with one line per register value
		list = make([]ls, 0, len(rv.NumberValues)+len(rv.TextValues)+len(rv.EnumValues))

		for _, v := range rv.NumberValues {
			list = append(list, ls{s: v.String(), sort: v.Register.Sort})
		}
		for _, v := range rv.TextValues {
			list = append(list, ls{s: v.String(), sort: v.Register.Sort})
		}
		for _, v := range rv.EnumValues {
			list = append(list, ls{s: v.String(), sort: v.Register.Sort})
		}

		// sort list by register sort order
		sort.SliceStable(list, func(i, j int) bool { return list[i].sort < list[j].sort })

	}

	time2 := time.Now()

	fmt.Printf("fetched %d registers, initialization took: %s, fetching took: %s\n",
		len(list),
		time1.Sub(time0),
		time2.Sub(time1),
	)

	// output list
	for _, v := range list {
		fmt.Printf("%s\n", v.s)
	}

}

type ls struct {
	s    string
	sort int
}
