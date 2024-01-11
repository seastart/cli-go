package main

import (
	"fmt"

	"github.com/seastart/cli-go"
)

func main() {

	app := cli.NewCliApp("app desc", &cli.Option{
		Name:  "env",
		Dft:   "dev",
		Usage: "envirioment",
	})
	// list
	app.AddCommand("list", "list command", func(cmd *cli.Command, remaincmds []string) (err error) {
		env, err := cmd.App().OptVal("env")
		if err != nil {
			return
		}
		page, err := cmd.OptVal("page")
		if err != nil {
			return
		}
		fmt.Printf("app env=%v\n", env)
		fmt.Printf("into list command page=%v\n", page)
		return
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
