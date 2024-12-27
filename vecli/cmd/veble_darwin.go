package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ErrUnimplemented = errors.New("veble on macOS is currently not implemented")

func runScan(_ *cobra.Command, _ []string) {
	fmt.Println(ErrUnimplemented)
	os.Exit(3)
}

func runDecode(cmd *cobra.Command, args []string) {
	fmt.Println(ErrUnimplemented)
	os.Exit(3)
}
