package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{Short: "Simple archiver"}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
