package main

import (
	"fmt"

	"github.com/seastart/cli-go"
)

func main() {

	app := cli.NewCliApp("app desc")
	// "" means the main app is one command
	app.AddCommand("", "default command", func(subcmds []string, options map[string]*cli.Option) {
		fmt.Printf("into default command start=%v\n", options["start"].GetVal())
	}, &cli.Option{
		Name:  "start",
		Dft:   0,
		Usage: "begin no",
	})
	app.Run()
	// ./main
	// ./main -start=2
}
