package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	cfg "gobn.github.io/rdf4go/internal/config"
	"gobn.github.io/rdf4go/internal/utils"
)

var (
	runCmd = &cobra.Command{
		Use:  "run",
		Args: cobra.ExactArgs(0),
		RunE: runCmdRunE,
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
}

func runCmdRunE(cmd *cobra.Command, args []string) error {
	utils.Flog().Info("entry: ", commandPath(cmd))
	protoConfig := cmd.Context().Value(ContextConfig)
	utils.Flog().Debugf("protoConfig = %s", protoConfig)
	config, isConfig := protoConfig.(*cfg.Config)
	if !isConfig {
		return errors.New("failed to obtain config from context")
	}
	utils.Flog().Infof("address = %s", config.Address())

	lis, err := net.Listen("tcp", config.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	conns := clientConns(lis)
	for {
		go handleConn(<-conns)
	}
}

func clientConns(lis net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := lis.Accept()
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		_, _ = client.Write(line)
	}
}
