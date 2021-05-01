package main

import (
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use: "config",
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
}
