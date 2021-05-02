// SPDX-FileCopyrightText: 2021 Iwan Aucamp
//
// SPDX-License-Identifier: CC0-1.0 OR Apache-2.0

package main

import (
	"context"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	cfg "gobn.github.io/rdf4go/internal/config"
	"gobn.github.io/rdf4go/internal/utils"
	"modernc.org/mathutil"
)

type key int

const (
	// ContextConfig is the context key for the configration.
	ContextConfig key = iota
)

func init() {
	var formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05.999", FullTimestamp: true}
	logrus.SetFormatter(formatter)
	levelString := os.Getenv("LOGRUS_LEVEL")
	useDefault := true
	if len(levelString) > 0 {
		setLevel, err := strconv.ParseUint(levelString, 10, 64)
		if err == nil {
			maxLevel := uint64(logrus.TraceLevel)
			minLevel := uint64(logrus.PanicLevel)
			var newLevel = mathutil.MaxUint64(mathutil.MinUint64(setLevel, maxLevel), minLevel)
			logrus.SetLevel(logrus.Level(newLevel))
		}
	}
	if useDefault {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func commandPath(cmd *cobra.Command) string {
	path := []string{}
	var cmdPtr *cobra.Command = cmd
	for cmdPtr != nil {
		path = append([]string{cmdPtr.Use}, path...)
		cmdPtr = cmdPtr.Parent()
	}
	return strings.Join(path, "/")
}

func main() {
	utils.Flog().Debugf("entry : args = %v, ...", os.Args)

	err := godotenv.Load(".env", "default.env")
	_, ok := err.(*fs.PathError)
	if err != nil && !ok {
		utils.Flog().Fatalf("Error loading .env files")
		panic(err)
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextConfig, cfg.NewConfig())
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		utils.Flog().Debugf("got err = %v", err)
		os.Exit(1)
	}
}
