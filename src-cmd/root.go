// Copyright © 2021 RackCloud

package src-cmd

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "rackcloud-cli",
	Short: "rackcloud-cli is a official command line interface for the RackCloud API"
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	    fmt.Fprintln(os.Stderr, err)
	    os.Exit(1)
	}
}

func init() {
    // Commands
	rootCmd.AddCommand(VersionCMD)
}

func initConfig() {
	
}