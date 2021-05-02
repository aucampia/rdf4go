// SPDX-FileCopyrightText: 2021 Iwan Aucamp
//
// SPDX-License-Identifier: CC0-1.0 OR Apache-2.0

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
