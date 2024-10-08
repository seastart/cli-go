package main

import (
	"fmt"

	"github.com/seastart/cli-go"
)

func main() {
	app := cli.NewCliApp("app desc")
	// test command with subcommand
	testCmd := app.AddCommandN("test", "test command with subcommand", nil, &cli.Option{
		Name: "start",
		Dft:  0,
		Desc: "begin no",
	})
	testCmd.SetPreRun(func(cmd *cli.Command, remaincmds []string) (err error) {
		start := cmd.OptVal("start")
		fmt.Printf("into test command prerun start=%v\n", start)
		return
	})
	subCmd := testCmd.AddCommandN("live", "test subcommand", func(cmd *cli.Command, remaincmds []string) (err error) {
		id := cmd.OptVal("id")
		fmt.Printf("into live subcommand run id=%v\n", id)
		return
	}, &cli.Option{
		Name: "id",
		Dft:  0,
		Desc: "live id",
	})
	subCmd.SetPreRun(func(cmd *cli.Command, remaincmds []string) (err error) {
		start := cmd.ParentCommand().OptVal("start")
		fmt.Printf("into live subcommand prerun start=%v\n", start)
		return
	})
	app.Run()
	// ./main
	// ./main test -start=2 live -id=100
	// ./main test -start=2 live -id=100 remaincmd1 remaincmd2
}
