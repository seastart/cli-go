package main

import (
	"fmt"

	"github.com/seastart/cli-go"
)

func main() {
	app := cli.NewCliApp("app desc")
	// test command
	app.AddCommand("test", "test command", func(cmd *cli.Command, remaincmds []string) (err error) {
		start := cmd.OptVal("start")
		fmt.Printf("into test command start=%v\n", start)
		return
	}, &cli.Option{
		Name: "start",
		Dft:  0,
		Desc: "begin no",
	})
	// list command
	app.AddCommand("list", "list command", func(cmd *cli.Command, remaincmds []string) (err error) {
		page := cmd.OptVal("page")
		fmt.Printf("into list command page=%v\n", page)
		return
	}, &cli.Option{
		Name: "page",
		Dft:  1,
		Desc: "page no",
	})
	app.Run()
	// ./main
	// ./main test -start=2
	// ./main list -page=3
}
