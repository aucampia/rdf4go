package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	cfg "gobn.github.io/rdf4go/internal/config"
	"gobn.github.io/rdf4go/internal/utils"
)

var (
	configDumpCmd = &cobra.Command{
		Use:  "dump",
		Args: cobra.ExactArgs(0),
		RunE: configDumpRunE,
	}
)

func init() {
	configCmd.AddCommand(configDumpCmd)
}

func configDumpRunE(cmd *cobra.Command, args []string) error {
	utils.Flog().Debugf("entry: %s", commandPath(cmd))
	protoConfig := cmd.Context().Value(ContextConfig)
	utils.Flog().Debugf("protoConfig = %s", protoConfig)
	config, isConfig := protoConfig.(*cfg.Config)
	if !isConfig || config == nil {
		return fmt.Errorf("failed to obtain config from context")
	}
	configs, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Printf("%s\n", configs)
	if err != nil {
		return err
	}
	return nil
}
