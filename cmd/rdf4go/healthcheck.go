package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	healthcheckCmd = &cobra.Command{
		Use:  "healthcheck",
		Args: cobra.ExactArgs(0),
		RunE: healthcheckRunE,
	}
)

func init() {
	rootCmd.AddCommand(healthcheckCmd)
}

func healthcheckRunE(cmd *cobra.Command, args []string) error {
	_, err := fmt.Printf("OK\n")
	if err != nil {
		return err
	}
	return nil
}
