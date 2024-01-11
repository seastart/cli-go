package main

import (
	"fmt"

	"github.com/seastart/cli-go"
)

func main() {
	app := cli.NewCliApp("app desc")
	// test command with subcommand
	testCmd := app.AddCommand("test", "test command with subcommand", nil, &cli.Option{
		Name:  "start",
		Dft:   0,
		Usage: "begin no",
	})
	testCmd.AddCommand("live", "test subcommand", func(cmd *cli.Command, remaincmds []string) (err error) {
		start, err := cmd.ParentCommand().OptVal("start")
		if err != nil {
			return
		}
		id, err := cmd.OptVal("id")
		if err != nil {
			return
		}
		fmt.Printf("into live subcommand parent start=%v\n", start)
		fmt.Printf("into live subcommand id=%v\n", id)
		return
	}, &cli.Option{
		Name:  "id",
		Dft:   0,
		Usage: "live id",
	})
	app.Run()
	// ./main
	// ./main test -start=2 live -id=100
	// ./main test -start=2 live -id=100 remaincmd1 remaincmd2
}
