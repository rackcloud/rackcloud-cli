// Copyright © 2021 RackCloud

package src-cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCMD = &cobra.Command{
	Use:   "version",
	Short: "Version number of RackCloud-CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println(RcVersion())
	},
}

func RcVersion() string {
	return "RackCloud-CLI v0.0.2"
}