package main

import (
	"fmt"

	"github.com/seastart/cli-go"
)

func main() {
	app := cli.NewCliApp("app desc")
	// test command with subcommand
	app.AddCommand("test", "test command with subcommand", func(subcmds []string, options map[string]*cli.Option) {
		if len(subcmds) == 0 {
			app.Exitf(1, "need specify subcommand: test who?")
		}
		fmt.Printf("into test command subcmds=%v start=%v\n", subcmds, options["start"].GetVal())
	}, &cli.Option{
		Name:  "start",
		Dft:   0,
		Usage: "begin no",
	})
	app.Run()
	// ./main
	// ./main test -start=2 live
}
