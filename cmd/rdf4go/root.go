// SPDX-FileCopyrightText: 2021 Iwan Aucamp
//
// SPDX-License-Identifier: CC0-1.0 OR Apache-2.0

package main

import (
	"errors"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	cfg "gobn.github.io/rdf4go/internal/config"
	"gobn.github.io/rdf4go/internal/utils"
	"modernc.org/mathutil"
)

var (
	rootCmd = &cobra.Command{
		Use:               path.Base(os.Args[0]),
		PersistentPreRunE: rootPersistentPreRunE,
		RunE:              rootRunE,
		SilenceUsage:      true,
	}
)

func init() {
	var rootPFlags = rootCmd.PersistentFlags()
	rootPFlags.CountP("verbose", "v", "counted verbosity")
}

func rootPersistentPreRunE(cmd *cobra.Command, args []string) error {
	verbosity, err := cmd.Flags().GetCount("verbose")
	if err != nil {
		return err
	}

	var setLevel = uint32(logrus.GetLevel()) + uint32(verbosity)
	var maxLevel = uint32(logrus.TraceLevel)
	var newLevel = mathutil.MinUint32(setLevel, maxLevel)

	logrus.SetLevel(logrus.Level(newLevel))
	utils.Flog().Debugf(
		"verbosity = %v, setLevel = %v, maxLevel = %v, newLevel = %v",
		verbosity, setLevel, maxLevel, newLevel)

	utils.Flog().Debugf("logrus.GetLevel() = %v", logrus.GetLevel())

	protoConfig := cmd.Context().Value(ContextConfig)
	utils.Flog().Debugf("protoConfig = %s", protoConfig)
	config, isConfig := protoConfig.(*cfg.Config)
	if !isConfig {
		return errors.New("failed to obtain config from context")
	}
	err = config.Load(nil)
	if err != nil {
		return err
	}

	return nil
}

func rootRunE(cmd *cobra.Command, args []string) error {
	utils.Flog().Debugf("entry: %v", cmd.Use)
	return nil
}
