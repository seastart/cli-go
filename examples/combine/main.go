package main

import (
	"fmt"

	"github.com/seastart/cli-go"
)

func main() {

	app := cli.NewCliApp("app desc")
	// set common options start
	app.AddCommand("", "default command", func(subcmds []string, options map[string]*cli.Option) {
		fmt.Printf("into default command env=%v\n", options["env"].GetVal())
	}, &cli.Option{
		Name:  "env",
		Dft:   "dev",
		Usage: "envirioment",
	})
	// list
	app.AddCommand("list", "list command", func(subcmds []string, options map[string]*cli.Option) {
		fmt.Printf("into list command page=%v\n", options["page"].GetVal())

	}, &cli.Option{
		Name:  "page",
		Dft:   1,
		Usage: "page no",
	})
	app.Run()
	// ./main
	// ./main -env=prod
	// ./main -env=prod list -page=3
}
