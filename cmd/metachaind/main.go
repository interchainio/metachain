package main

import (
	"os"

	"github.com/interchainberlin/metachain/cmd/metachaind/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
