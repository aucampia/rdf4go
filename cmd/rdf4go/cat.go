// SPDX-FileCopyrightText: 2021 Iwan Aucamp
//
// SPDX-License-Identifier: CC0-1.0 OR Apache-2.0

package main

import (
	"github.com/spf13/cobra"
	"gobn.github.io/rdf4go/internal/utils"
)

var (
	runCmd = &cobra.Command{
		Use:  "cat",
		Args: cobra.MinimumNArgs(0),
		RunE: catRunE,
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
}

func catRunE(cmd *cobra.Command, args []string) error {
	utils.Flog().Info("entry: ", commandPath(cmd))
	return nil
}
