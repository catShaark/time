package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/romelukaku/time/app"
	"github.com/romelukaku/time/cmd/timed/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
