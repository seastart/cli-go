package main

import (
	"fmt"

	"github.com/seastart/cli-go"
)

func main() {
	app := cli.NewCliApp("app desc")
	// test command
	app.AddCommand("test", "test command", func(subcmds []string, options map[string]*cli.Option) {
		fmt.Printf("into test command start=%v\n", options["start"].GetVal())

	}, &cli.Option{
		Name:  "start",
		Dft:   0,
		Usage: "begin no",
	})
	// list command
	app.AddCommand("list", "list command", func(subcmds []string, options map[string]*cli.Option) {
		fmt.Printf("into list command page=%v\n", options["page"].GetVal())

	}, &cli.Option{
		Name:  "page",
		Dft:   1,
		Usage: "page no",
	})
	app.Run()
	// ./main
	// ./main test -start=2
	// ./main list -page=3
}
