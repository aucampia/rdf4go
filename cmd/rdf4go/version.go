package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:  "version",
		Args: cobra.ExactArgs(0),
		RunE: versionRunE,
	}
	gitDesc   string
	gitCommit string
	gitRemote string
	buildDT   string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func versionRunE(cmd *cobra.Command, args []string) error {
	info := map[string]string{
		"executable":      path.Base(os.Args[0]),
		"git_description": gitDesc,
		"git_commit":      gitCommit,
		"git_remote":      gitRemote,
		"build_datetime":  buildDT,
	}
	infos, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	_, err = fmt.Printf("%s\n", infos)
	if err != nil {
		return err
	}
	return nil
}
